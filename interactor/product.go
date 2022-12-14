package interactor

import (
	"context"
	"go-ecommerce/adapter/db"
	"go-ecommerce/dto"
)

func createProductResponse(product db.Product) dto.ProductResponse {

	var productVariations []dto.ProductVariation

	for _, productVariation := range product.ProductVariations {
		productVariations = append(productVariations, dto.ProductVariation{
			ID:    productVariation.ID,
			Name:  productVariation.Name,
			Price: uint(productVariation.Price),
			Stock: uint(productVariation.Stock),
		})
	}

	return dto.ProductResponse{
		ID:                product.ID,
		Name:              product.Name,
		ProductVariations: productVariations,
		ProductCategory: dto.ProductCategoryResponse{
			ID:          product.ProductCategory.ID,
			Name:        product.ProductCategory.Name,
			Description: product.ProductCategory.Description,
		},
	}

}

func (i Interactor) IndexProduct(ctx context.Context) (dto.IndexProductResponse, error) {

	products, err := i.store.IndexProducts()

	if err != nil {
		return dto.IndexProductResponse{}, err
	}

	var productsResponse []dto.ProductResponse
	for _, product := range products {

		productsResponse = append(productsResponse, createProductResponse(product))

	}
	return dto.IndexProductResponse{Products: productsResponse}, nil
}

func (i Interactor) FindProduct(productId uint) (dto.ProductResponse, error) {

	product, err := i.store.FindProduct(productId)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return createProductResponse(product), nil
}

func (i Interactor) StoreProduct(ctx context.Context, req dto.StoreProductRequest) (dto.ProductResponse, error) {

	product := db.Product{
		Name:              req.Name,
		ProductCategoryID: req.ProductCategoryId,
		Description:       req.Description,
	}

	var variations []db.ProductVariation

	for _, variation := range req.Variations {
		variations = append(variations, db.ProductVariation{
			Name:  variation.Name,
			Price: variation.Price,
			Stock: variation.Stock,
		})
	}

	product.ProductVariations = variations

	product, err := i.store.StoreProduct(product)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return createProductResponse(product), nil
}

func (i Interactor) DestroyProduct(ctx context.Context, productId uint) error {
	err := i.store.DestroyProduct(productId)
	if err != nil {
		return err
	}
	return nil
}
