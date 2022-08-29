package contract

import (
	"context"
	"go-ecommerce/adapter/db"
	"time"
)

type Repository interface {
	RegisterUser(ctx context.Context, user db.User) (db.User, error)
	LoginUser(email string, cellNumber string) (db.User, error)
	GenerateClientAccessToken(userId uint, expiryData time.Time, title string, token string) (db.AccessToken, error)

	StoreProductCategory(productCategory db.ProductCategory) (db.ProductCategory, error)
	IndexProductCategories() ([]db.ProductCategory, error)
	GetProductCategory(productCategoryId uint) (db.ProductCategory, error)

	IndexProducts() ([]db.Product, error)
	FindProduct(productId uint) (db.Product, error)
	StoreProduct(product db.Product) (db.Product, error)
	DestroyProduct(productId uint) error
}
