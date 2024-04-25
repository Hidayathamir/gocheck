package main

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// printDDL print DDL of table.
func printDDL(table any) {
	cfg, err := config.Load("../../../config.yml")
	fatalIfErr(err)

	pg, err := db.NewPostgres(*cfg, db.WithGormConfig(&gorm.Config{DryRun: true}))
	fatalIfErr(err)

	err = pg.DB.AutoMigrate(table)
	fatalIfErr(err)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
