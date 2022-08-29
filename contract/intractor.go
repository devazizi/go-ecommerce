package contract

import (
	"context"
	"go-ecommerce/dto"
)

type Intractor interface {
	RegisterUser(ctx context.Context, request dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	LoginUser(ctx context.Context, request dto.LoginUserRequest) (dto.LoginUserResponse, error)

	StoreProductCategory(ctx context.Context, req dto.StorePostCategoryRequest) (dto.ProductCategoryResponse, error)
	IndexProductCategory(ctx context.Context) (dto.IndexProductCategoriesResponse, error)
	FindProductCategory(ctx context.Context, productCategoryId uint) (dto.ProductCategoryResponse, error)

	IndexProduct(ctx context.Context) (dto.IndexProductResponse, error)
	FindProduct(productId uint) (dto.ProductResponse, error)
	StoreProduct(ctx context.Context, req dto.StoreProductRequest) (dto.ProductResponse, error)
	DestroyProduct(ctx context.Context, productId uint) error
}
