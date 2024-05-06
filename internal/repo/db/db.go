package db

import (
	"fmt"
	"time"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/Hidayathamir/txmanager"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres represents a PostgreSQL database connection.
type Postgres struct {
	DB        *gorm.DB
	TxManager txmanager.ITransactionManager
}

// NewPostgresOpt contains options for configuring the PostgreSQL connection.
type NewPostgresOpt struct {
	gormConfig *gorm.Config
}

// NewPostgresOption defines a function signature for configuring NewPostgresOpt.
type NewPostgresOption func(*NewPostgresOpt)

// NewPostgres initializes a new PostgreSQL database connection.
func NewPostgres(cfg *config.Config, options ...NewPostgresOption) (*Postgres, error) {
	option := &NewPostgresOpt{
		gormConfig: &gorm.Config{},
	}
	for _, opt := range options {
		opt(option)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.GetPostgresHost(), cfg.GetPostgresUsername(), cfg.GetPostgresPassword(),
		cfg.GetPostgresDBName(), cfg.GetPostgresPort(),
	)

	var db *gorm.DB
	var errInitDB error
	const maxAttemptInitDB = 10
	for i := 0; i < maxAttemptInitDB; i++ {
		db, errInitDB = gorm.Open(postgres.Open(dsn), option.gormConfig)
		if errInitDB == nil {
			break
		}

		errInitDB = fmt.Errorf("error initialize db session: %w", errInitDB)

		logrus.
			WithField("attempt left", maxAttemptInitDB-i-1).
			Warn(trace.Wrap(errInitDB))

		time.Sleep(time.Second)
	}
	if errInitDB != nil {
		errInitDB := fmt.Errorf("error initialize db session %d times: %w", maxAttemptInitDB, errInitDB)
		return nil, trace.Wrap(errInitDB)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	sqlDB.SetMaxIdleConns(cfg.GetGormMaxIdleConns())
	sqlDB.SetMaxOpenConns(cfg.GetGormMaxOpenConns())

	txManager := txmanager.NewTransactionManager(db)

	pg := &Postgres{
		DB:        db,
		TxManager: txManager,
	}

	logrus.Info("success create db connection 🟢")

	return pg, nil
}

// WithGormConfig configures the GORM options.
func WithGormConfig(gormConfig *gorm.Config) NewPostgresOption {
	return func(npo *NewPostgresOpt) {
		npo.gormConfig = gormConfig
	}
}
