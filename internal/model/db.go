package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Movie{},
		&UserMovie{},
		&Quote{},
		&UserTokens{},
	)
}
