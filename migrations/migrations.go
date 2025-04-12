package migrations

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/404th/smtest/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Migrator struct {
	config *config.Config
	logger *zap.Logger
}

func NewMigrator(config *config.Config, logger *zap.Logger) *Migrator {
	return &Migrator{
		config: config,
		logger: logger,
	}
}

func (m *Migrator) RunMigrations() error {
	m.logger.Info("Checking for new SQL migrations", zap.String("migration_path", m.config.MigrationPath))

	// Validate and create migration directory if missing
	path := strings.TrimPrefix(m.config.MigrationPath, "file://")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		m.logger.Warn("Migration directory does not exist, creating", zap.String("path", path))
		if err := os.MkdirAll(path, 0755); err != nil {
			m.logger.Error("Failed to create migration directory", zap.Error(err))
			return fmt.Errorf("failed to create migration directory %s: %w", path, err)
		}
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=%s",
		m.config.PostgresUsername,
		m.config.PostgresPassword,
		m.config.PostgresHost,
		m.config.PostgresPort,
		m.config.PostgresName,
		m.config.PostgresTimezone,
	)
	m.logger.Debug("Using DSN", zap.String("dsn", dsn))

	mig, err := migrate.New(m.config.MigrationPath, dsn)
	if err != nil {
		m.logger.Error("Failed to initialize migrator", zap.Error(err))
		return err
	}

	err = mig.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			m.logger.Info("No new migrations to apply")
			return nil
		}
		m.logger.Error("Failed to apply migrations", zap.Error(err))
		return err
	}

	m.logger.Info("New SQL migrations applied successfully")
	return nil
}

var Module = fx.Options(
	fx.Provide(NewMigrator),
	fx.Invoke(func(lc fx.Lifecycle, migrator *Migrator) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return migrator.RunMigrations()
			},
			OnStop: func(ctx context.Context) error {
				migrator.logger.Info("App stopped")
				return nil
			},
		})
	}),
)
