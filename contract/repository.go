package contract

import (
	"context"
	"go-ecommerce/adapter/db"
)

type Repository interface {
	// auth
	RegisterUser(ctx context.Context, user db.User) (db.User, error)
	LoginUser(ctx context.Context, user db.User) (db.User, error)

	// product category
	StoreProductCategory(productCategory db.ProductCategory) (db.ProductCategory, error)
	IndexProductCategories() ([]db.ProductCategory, error)
	GetProductCategory(productCategoryId uint) (db.ProductCategory, error)
}
