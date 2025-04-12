package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/404th/smtest/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDatabase(dsn, dbname string) error {
	createDSN := strings.Replace(dsn, "dbname="+dbname, "dbname=postgres", 1)
	db, err := sql.Open("postgres", createDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = db.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewGormDB(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.PostgresHost,
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresName,
		config.PostgresPort,
		config.PostgresTimezone,
	)

	err := createDatabase(dsn, config.PostgresName)
	if err != nil {
		logger.Error("Failed to create database", zap.Error(err))
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to open GORM connection", zap.Error(err))
		return nil, err
	}
	return db, nil
}

var Module = fx.Options(
	fx.Provide(NewGormDB),
	fx.Invoke(func(lc fx.Lifecycle, db *gorm.DB, logger *zap.Logger) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("GORM connection established")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info("Closing GORM connection")
				sqlDB, err := db.DB()
				if err != nil {
					logger.Error("Failed to get sql.DB", zap.Error(err))
					return err
				}
				return sqlDB.Close()
			},
		})
	}),
)
