// Package config -.
package config

import (
	"fmt"
	"strings"

	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/spf13/viper"
)

// Config -.
type Config struct{ *viper.Viper }

// Load loads config.yml and env var.
func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	cfg := &Config{Viper: v}

	err = cfg.Validate()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return cfg, nil
}

// Validate -.
func (c *Config) Validate() error {
	switch c.GetAppEnvironment() {
	case "dev", "prod":
	default:
		err := fmt.Errorf("unknown app environment '%s'", c.GetAppEnvironment())
		return trace.Wrap(err)
	}

	switch c.GetLoggerLogLevel() {
	case "panic", "fatal", "error", "warn", "warning", "info", "debug", "trace":
	default:
		err := fmt.Errorf("unknown logger log level '%s'", c.GetLoggerLogLevel())
		return trace.Wrap(err)
	}

	return nil
}

// GetAppName -.
func (c *Config) GetAppName() string {
	return c.GetString("app.name")
}

// GetAppVersion -.
func (c *Config) GetAppVersion() string {
	return c.GetString("app.version")
}

// GetAppEnvironment -.
func (c *Config) GetAppEnvironment() string {
	return c.GetString("app.environment")
}

// GetHTTPHost -.
func (c *Config) GetHTTPHost() string {
	return c.GetString("http.host")
}

// GetHTTPPort -.
func (c *Config) GetHTTPPort() string {
	return c.GetString("http.port")
}

// GetGRPCHost -.
func (c *Config) GetGRPCHost() string {
	return c.GetString("grpc.host")
}

// GetGRPCPort -.
func (c *Config) GetGRPCPort() string {
	return c.GetString("grpc.port")
}

// GetLoggerLogLevel -.
func (c *Config) GetLoggerLogLevel() string {
	return c.GetString("logger.log_level")
}

// GetPostgresUsername -.
func (c *Config) GetPostgresUsername() string {
	return c.GetString("postgres.username")
}

// GetPostgresPassword -.
func (c *Config) GetPostgresPassword() string {
	return c.GetString("postgres.password")
}

// GetPostgresHost -.
func (c *Config) GetPostgresHost() string {
	return c.GetString("postgres.host")
}

// GetPostgresPort -.
func (c *Config) GetPostgresPort() string {
	return c.GetString("postgres.port")
}

// GetPostgresDBName -.
func (c *Config) GetPostgresDBName() string {
	return c.GetString("postgres.db_name")
}

// GetRedisHost -.
func (c *Config) GetRedisHost() string {
	return c.GetString("redis.host")
}

// GetRedisPort -.
func (c *Config) GetRedisPort() string {
	return c.GetString("redis.port")
}

// GetMigrationAuto -.
func (c *Config) GetMigrationAuto() bool {
	return c.GetBool("migration.auto")
}

// GetMigrationRequired -.
func (c *Config) GetMigrationRequired() bool {
	return c.GetBool("migration.required")
}
