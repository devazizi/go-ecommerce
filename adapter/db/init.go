package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	store *gorm.DB
}

func NewDB(dsn string) DB {

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to db")
	}

	var entities = []any{
		&User{},
		&Product{},
		&ProductVariation{},
		&ShoppingCart{},
		&ShoppingCartItem{},
		&Order{},
		&OrderItem{},
		&Refund{},
		&RefundItem{},
		&PostCategory{},
		&Post{},
		&Comment{},
		&Role{},
		&Permission{},
		&Province{},
		&City{},
		&Address{},
	}

	if migrationErr := database.AutoMigrate(entities...); migrationErr != nil {
		panic("db migration fail")
	}

	return DB{store: database}
}
