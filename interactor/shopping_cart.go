package interactor

import (
	"context"
	"go-ecommerce/dto"
)

func (i Interactor) GetShoppingCart(ctx context.Context) (dto.ShoppingCartResponse, error) {
	shoppingCart := i.store.GetShoppingCart(1)

	var shoppingCartItemsResponse []dto.ShoppingCartItemResponse

	for _, shoppingCartItem := range shoppingCart.ShoppingCartItems {

		shoppingCartItemsResponse = append(shoppingCartItemsResponse, dto.ShoppingCartItemResponse{
			ID:       shoppingCartItem.ID,
			Quantity: shoppingCartItem.Quantity,
			ProductVariation: dto.ProductVariation{
				ID:    shoppingCartItem.ProductVariation.ID,
				Name:  shoppingCartItem.ProductVariation.Name,
				Stock: shoppingCartItem.ProductVariation.Stock,
				Price: shoppingCartItem.ProductVariation.Price,
				Product: dto.ProductResponse{
					ID:   shoppingCartItem.ProductVariation.Product.ID,
					Name: shoppingCartItem.ProductVariation.Product.Name,
					ProductCategory: dto.ProductCategoryResponse{
						ID:   shoppingCartItem.ProductVariation.Product.ProductCategory.ID,
						Name: shoppingCartItem.ProductVariation.Product.ProductCategory.Name,
					},
				},
			},
		})
	}

	return dto.ShoppingCartResponse{
		ID:                shoppingCart.ID,
		ShoppingCartItems: shoppingCartItemsResponse,
	}, nil
}
