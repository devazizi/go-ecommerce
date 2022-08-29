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

func (db DB) StoreProduct(product Product) (Product, error) {
	db.store.Create(&product)

	db.store.
		Preload("ProductVariations").
		Preload("ProductCategory").
		First(&product, product.ID)

	return product, nil
}

func (db DB) DestroyProduct(productId uint) error {

	err := db.store.Transaction(func(gorm *gorm.DB) error {
		db.store.Where("product_id = ?", productId).Delete(&ProductVariation{})
		db.store.Delete(&Product{}, productId)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
