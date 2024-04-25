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

// Up -.
func Up(db *gorm.DB) error {
	migrate.SetTable("migrations")

	var countMigrationApplied int
	var errMigrateUp error
	maxAttemptMigrateUp := 10
	for i := 0; i < maxAttemptMigrateUp; i++ {
		sql, err := db.DB()
		if err != nil {
			return trace.Wrap(err)
		}

		countMigrationApplied, errMigrateUp = migrate.Exec(sql, "postgres", getFileMigrationSource(), migrate.Up)
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

func getFileMigrationSource() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join("internal", "table", "migration", "schema_migration"),
	}
	return migrations
}
