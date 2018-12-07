package mysql

import (
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/jinzhu/gorm"
)

// https://github.com/go-gormigrate/gormigrate
func Migrate(d *gorm.DB) {
	d.AutoMigrate(
		db.User{},
		db.Item{},
		db.Account{},
	)

	// TODO: move to a better migration tool: https://github.com/go-gormigrate/gormigrate
}
