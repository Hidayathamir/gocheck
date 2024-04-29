// Package migrate -.
package migrate

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Hidayathamir/gocheck/pkg/trace"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// UpOpt -.
type UpOpt struct {
	Dir string
}

// UpOption -.
type UpOption func(*UpOpt)

var defaultDir = filepath.Join("internal", "repo", "db", "migration", "schema_migration")

// Up -.
func Up(db *gorm.DB, options ...UpOption) error {
	option := &UpOpt{Dir: defaultDir}
	for _, opt := range options {
		opt(option)
	}

	migrate.SetTable("migrations")

	var countMigrationApplied int
	var errMigrateUp error
	maxAttemptMigrateUp := 10
	for i := 0; i < maxAttemptMigrateUp; i++ {
		sql, err := db.DB()
		if err != nil {
			return trace.Wrap(err)
		}

		fileMigrationSource := &migrate.FileMigrationSource{Dir: option.Dir}
		countMigrationApplied, errMigrateUp = migrate.Exec(sql, "postgres", fileMigrationSource, migrate.Up)
		if errMigrateUp == nil {
			break
		}

		errMigrateUp = fmt.Errorf("error migrate up: %w", errMigrateUp)

		logrus.
			WithField("attempt_left", maxAttemptMigrateUp-i-1).
			Warn(trace.Wrap(errMigrateUp))

		time.Sleep(time.Second)
	}
	if errMigrateUp != nil {
		errMigrateUp := fmt.Errorf("error migrate up %d times: %w", maxAttemptMigrateUp, errMigrateUp)
		return trace.Wrap(errMigrateUp)
	}

	logrus.Infof("migration success, %d applied 🟢", countMigrationApplied)

	return nil
}

// WithDir -.
func WithDir(dir string) UpOption {
	return func(uo *UpOpt) {
		uo.Dir = dir
	}
}
