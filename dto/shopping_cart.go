package dto

type ShoppingCartItemResponse struct {
	ID               uint             `json:"id"`
	ProductVariation ProductVariation `json:"product_variation"`
	Quantity         uint             `json:"quantity"`
}

type ShoppingCartResponse struct {
	ID                uint                       `json:"id"`
	ShoppingCartItems []ShoppingCartItemResponse `json:"shopping_cart_items"`
}
