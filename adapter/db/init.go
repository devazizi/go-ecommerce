package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(dsn string) DB {

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to db")
	}

	var entities []any

	if migrationErr := database.AutoMigrate(entities...); migrationErr != nil {
		panic("db migration fail")
	}

	return DB{db: database}
}
