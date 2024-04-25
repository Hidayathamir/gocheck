// Package db -.
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

// Postgres -.
type Postgres struct {
	DB        *gorm.DB
	TxManager txmanager.ITransactionManager
}

// NewPostgresOpt -.
type NewPostgresOpt struct {
	gormConfig *gorm.Config
}

// NewPostgresOption -.
type NewPostgresOption func(*NewPostgresOpt)

// NewPostgres -.
func NewPostgres(cfg config.Config, options ...NewPostgresOption) (*Postgres, error) {
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

	txManager := txmanager.NewTransactionManager(db)

	pg := &Postgres{
		DB:        db,
		TxManager: txManager,
	}

	logrus.Info("success create db connection 🟢")

	return pg, nil
}

// WithGormConfig -.
func WithGormConfig(gormConfig *gorm.Config) NewPostgresOption {
	return func(npo *NewPostgresOpt) {
		npo.gormConfig = gormConfig
	}
}
