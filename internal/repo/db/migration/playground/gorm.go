package main

import (
	"path/filepath"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// gormPlayground print DDL of table.
func gormPlayground(fn func(pg *gorm.DB)) {
	cfg, err := config.Load(filepath.Join("..", "..", "..", "..", "config.yml"))
	fatalIfErr(err)

	gormConfig := &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Info),
	}

	pg, err := db.NewPostgres(*cfg, db.WithGormConfig(gormConfig))
	fatalIfErr(err)

	fn(pg.DB)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
