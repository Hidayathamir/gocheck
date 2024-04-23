// Package app contains application starter.
package app

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/pkg/trace"
	"github.com/sirupsen/logrus"
)

// Run runs application.
func Run() {
	cfg, err := config.Load()
	fatalIfErr(err)

	logrus.Info(cfg.AllSettings())
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
