package db

import (
	"errors"
	"gorm.io/gorm"
)

func (db DB) IndexProducts() ([]Product, error) {
	var products []Product

	db.store.Preload("ProductVariations").Preload("ProductCategory").Find(&products)

	return products, nil
}

func (db DB) FindProduct(productId uint) (Product, error) {
	var product Product
	result := db.store.
		Preload("ProductVariations").
		Preload("ProductCategory").
		Where("id = ?", productId).
		First(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return product, result.Error
	}

	return product, nil
}
