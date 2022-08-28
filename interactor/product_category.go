package interactor

import (
	"context"
	"go-ecommerce/adapter/db"
	"go-ecommerce/dto"
)

func (i Interactor) StoreProductCategory(ctx context.Context, req dto.StorePostCategoryRequest) (dto.ProductCategoryResponse, error) {

	productCategory, err := i.store.StoreProductCategory(db.ProductCategory{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return dto.ProductCategoryResponse{}, err
	}

	return dto.ProductCategoryResponse{
		ID:          productCategory.ID,
		Name:        productCategory.Name,
		Description: productCategory.Description,
	}, nil
}

func (i Interactor) IndexProductCategory(ctx context.Context) (dto.IndexProductCategoriesResponse, error) {
	productCategories, _ := i.store.IndexProductCategories()

	var productCategoriesResponse []dto.ProductCategoryResponse

	for _, productCategory := range productCategories {
		productCategoriesResponse = append(productCategoriesResponse, dto.ProductCategoryResponse{
			ID:          productCategory.ID,
			Name:        productCategory.Name,
			Description: productCategory.Description,
		})
	}

	return dto.IndexProductCategoriesResponse{ProductCategories: productCategoriesResponse}, nil
}

func (i Interactor) FindProductCategory(ctx context.Context, productCategoryId uint) (dto.ProductCategoryResponse, error) {

	productCategory, err := i.store.GetProductCategory(productCategoryId)
	if err != nil {
		return dto.ProductCategoryResponse{}, err
	}

	return dto.ProductCategoryResponse{
		ID:          productCategory.ID,
		Name:        productCategory.Name,
		Description: productCategory.Description,
	}, nil
}
