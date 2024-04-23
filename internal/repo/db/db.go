// Package db -.
package db

import (
	"fmt"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/pkg/trace"
	"github.com/Hidayathamir/txmanager"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres -.
type Postgres struct {
	DB        *gorm.DB
	TxManager txmanager.ITransactionManager
}

// NewPGPoolConn -.
func NewPGPoolConn(cfg config.Config) (*Postgres, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.GetPostgresHost(), cfg.GetPostgresUsername(), cfg.GetPostgresPassword(),
		cfg.GetPostgresDBName(), cfg.GetPostgresPort(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	txManager := txmanager.NewTransactionManager(db)

	pg := &Postgres{
		DB:        db,
		TxManager: txManager,
	}

	return pg, nil
}
