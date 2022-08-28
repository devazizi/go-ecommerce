package interactor

import (
	"context"
	"go-ecommerce/dto"
)

func (i Interactor) IndexProduct(ctx context.Context) (dto.IndexProductResponse, error) {

	products, err := i.store.IndexProducts()

	if err != nil {
		return dto.IndexProductResponse{}, err
	}

	var productsResponse []dto.ProductResponse
	for _, product := range products {

		var productVariations []dto.ProductVariation

		for _, productVariation := range product.ProductVariations {
			productVariations = append(productVariations, dto.ProductVariation{
				ID:    productVariation.ID,
				Name:  productVariation.Name,
				Price: uint(productVariation.Price),
				Stock: uint(productVariation.Stock),
			})
		}

		productResponse := dto.ProductResponse{
			ID:                product.ID,
			Name:              product.Name,
			ProductVariations: productVariations,
			ProductCategory: dto.ProductCategoryResponse{
				ID:          product.ProductCategory.ID,
				Name:        product.ProductCategory.Name,
				Description: product.ProductCategory.Description,
			},
		}

		productsResponse = append(productsResponse, productResponse)

	}
	return dto.IndexProductResponse{Products: productsResponse}, nil
}
