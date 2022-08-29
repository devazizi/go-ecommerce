package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (db DB) StoreProductCategory(productCategory ProductCategory) (ProductCategory, error) {
	db.store.Create(&productCategory)

	return productCategory, nil
}

func (db DB) IndexProductCategories() ([]ProductCategory, error) {
	var productCategories []ProductCategory
	db.store.Find(&productCategories)

	return productCategories, nil
}

func (db DB) GetProductCategory(productCategoryId uint) (ProductCategory, error) {
	var productCategory ProductCategory
	result := db.store.Where("id = ?", productCategoryId).First(&productCategory)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return productCategory, errors.New(fmt.Sprintf("product category with id %d not found.", productCategoryId))
	}

	return productCategory, nil
}
