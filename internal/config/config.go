package config

import (
	"fmt"
	"strings"

	"github.com/Hidayathamir/gocheck/pkg/errutil"
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct{ *viper.Viper }

// Load loads configuration settings from the specified file and environment variables.
func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, errutil.Wrap(err)
	}

	cfg := &Config{Viper: v}

	err = cfg.Validate()
	if err != nil {
		return nil, errutil.Wrap(err)
	}

	return cfg, nil
}

// Validate checks if the loaded configuration is valid.
func (c *Config) Validate() error {
	switch c.GetAppEnvironment() {
	case "dev", "prod":
	default:
		err := fmt.Errorf("unknown app environment '%s'", c.GetAppEnvironment())
		return errutil.Wrap(err)
	}

	switch c.GetLoggerLogLevel() {
	case "panic", "fatal", "error", "warn", "warning", "info", "debug", "trace":
	default:
		err := fmt.Errorf("unknown logger log level '%s'", c.GetLoggerLogLevel())
		return errutil.Wrap(err)
	}

	return nil
}

// GetAppName retrieves the application name from the configuration.
func (c *Config) GetAppName() string {
	return c.GetString("app.name")
}

// GetAppVersion retrieves the application version from the configuration.
func (c *Config) GetAppVersion() string {
	return c.GetString("app.version")
}

// GetAppEnvironment retrieves the application environment from the configuration.
func (c *Config) GetAppEnvironment() string {
	return c.GetString("app.environment")
}

// GetHTTPHost retrieves the HTTP host from the configuration.
func (c *Config) GetHTTPHost() string {
	return c.GetString("http.host")
}

// GetHTTPPort retrieves the HTTP port from the configuration.
func (c *Config) GetHTTPPort() string {
	return c.GetString("http.port")
}

// GetGRPCHost retrieves the gRPC host from the configuration.
func (c *Config) GetGRPCHost() string {
	return c.GetString("grpc.host")
}

// GetGRPCPort retrieves the gRPC port from the configuration.
func (c *Config) GetGRPCPort() string {
	return c.GetString("grpc.port")
}

// GetLoggerLogLevel retrieves the logger log level from the configuration.
func (c *Config) GetLoggerLogLevel() string {
	return c.GetString("logger.log_level")
}

// GetPostgresUsername retrieves the PostgreSQL username from the configuration.
func (c *Config) GetPostgresUsername() string {
	return c.GetString("postgres.username")
}

// GetPostgresPassword retrieves the PostgreSQL password from the configuration.
func (c *Config) GetPostgresPassword() string {
	return c.GetString("postgres.password")
}

// GetPostgresHost retrieves the PostgreSQL host from the configuration.
func (c *Config) GetPostgresHost() string {
	return c.GetString("postgres.host")
}

// GetPostgresPort retrieves the PostgreSQL port from the configuration.
func (c *Config) GetPostgresPort() string {
	return c.GetString("postgres.port")
}

// SetPostgresPort sets the PostgreSQL port in the configuration.
func (c *Config) SetPostgresPort(port string) {
	c.Set("postgres.port", port)
}

// GetPostgresDBName retrieves the PostgreSQL database name from the configuration.
func (c *Config) GetPostgresDBName() string {
	return c.GetString("postgres.db_name")
}

// GetRedisHost retrieves the Redis host from the configuration.
func (c *Config) GetRedisHost() string {
	return c.GetString("redis.host")
}

// GetRedisPort retrieves the Redis port from the configuration.
func (c *Config) GetRedisPort() string {
	return c.GetString("redis.port")
}

// SetRedisPort sets the Redis port in the configuration.
func (c *Config) SetRedisPort(port string) {
	c.Set("redis.port", port)
}

// GetMigrationAuto retrieves the migration auto setting from the configuration.
func (c *Config) GetMigrationAuto() bool {
	return c.GetBool("migration.auto")
}

// GetMigrationRequired retrieves the migration required setting from the configuration.
func (c *Config) GetMigrationRequired() bool {
	return c.GetBool("migration.required")
}

// GetGormMaxIdleConns retrieves the gorm maximum idle connection pooling from the configuration.
func (c *Config) GetGormMaxIdleConns() int {
	return c.GetInt("gorm.max_idle_conns")
}

// GetGormMaxOpenConns retrieves the gorm maximum open connection pooling from the configuration.
func (c *Config) GetGormMaxOpenConns() int {
	return c.GetInt("gorm.max_open_conns")
}
