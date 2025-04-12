package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AppName       string `json:"appName"`
	AppPort       string `json:"appPort"`
	AppHost       string `json:"appHost"`
	AppHTTPScheme string `json:"appHTTPScheme"`
	AppVersion    string `json:"appVersion"`

	PostgresName          string `json:"postgresName"`
	PostgresUsername      string `json:"postgresUsername"`
	PostgresPassword      string `json:"postgresPassword"`
	PostgresPort          string `json:"postgresPort"`
	PostgresHost          string `json:"postgresHost"`
	PostgresMaxConnection uint   `json:"postgresMaxConnection"`
	PostgresTimezone      string `json:"postgresTimezone"`

	JWTSecretSalt string `json:"jwtSecretSalt"`
	PasswordSalt  string `json:"passwordSalt"`
	MigrationPath string `json:"migrationPath"`
}

func NewConfig() (cfg *Config) {
	cfg = &Config{}
	_ = godotenv.Load("./.env")

	// Get absolute path for migrations
	wd, err := os.Getwd()
	if err != nil {
		return cfg
	}

	migrationPath := "file://" + filepath.Join(wd, "migrations", "postgres")

	cfg.AppName = cast.ToString(getEnvOrSetDefault("APP_NAME", "smtest"))
	cfg.AppPort = cast.ToString(getEnvOrSetDefault("APP_PORT", "3000"))
	cfg.AppHost = cast.ToString(getEnvOrSetDefault("APP_HOST", "localhost"))
	cfg.AppHTTPScheme = cast.ToString(getEnvOrSetDefault("HTTP_SCHEME", "http"))
	cfg.AppVersion = cast.ToString(getEnvOrSetDefault("APP_VERSION", "0.0.1"))

	cfg.PostgresName = cast.ToString(getEnvOrSetDefault("POSTGRES_NAME", "postgres"))
	cfg.PostgresUsername = cast.ToString(getEnvOrSetDefault("POSTGRES_USERNAME", "postgres"))
	cfg.PostgresPassword = cast.ToString(getEnvOrSetDefault("POSTGRES_PASSWORD", "postgres"))
	cfg.PostgresPort = cast.ToString(getEnvOrSetDefault("POSTGRES_PORT", "5432"))
	cfg.PostgresHost = cast.ToString(getEnvOrSetDefault("POSTGRES_HOST", "0.0.0.0"))
	cfg.PostgresMaxConnection = cast.ToUint(getEnvOrSetDefault("POSTGRES_MAX_CONNECTION", 30))
	cfg.PostgresTimezone = cast.ToString(getEnvOrSetDefault("POSTGRES_TIMEZONE", "Asia/Tashkent"))

	cfg.JWTSecretSalt = cast.ToString(getEnvOrSetDefault("JWT_SECRET_SALT", "supersecretstringhere"))
	cfg.PasswordSalt = cast.ToString(getEnvOrSetDefault("PASSWORD_SALT", "supersecretstring2here"))
	cfg.MigrationPath = cast.ToString(getEnvOrSetDefault("MIGRATION_PATH", migrationPath))

	return
}

func getEnvOrSetDefault(key string, defaultValue any) any {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
