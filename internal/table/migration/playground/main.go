// Package main -.
package main

import (
	"github.com/Hidayathamir/gocheck/internal/table"
	"gorm.io/gorm"
)

//nolint:gomnd
func main() {
	gormPlayground(func(pg *gorm.DB) {
		var err error

		err = pg.Create(&table.User{
			Username: "hidayat",
			FullName: "hidayat hamir",
			Balance:  200000,
		}).Error
		fatalIfErr(err)

		err = pg.Create(&table.User{
			Username: "hafiz",
			FullName: "hafiz arrahman",
			Balance:  700000,
		}).Error
		fatalIfErr(err)

		err = pg.Create(&table.User{
			Username: "aji",
			FullName: "aji hidayat",
			Balance:  400000,
		}).Error
		fatalIfErr(err)
	})
}
