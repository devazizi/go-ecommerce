package dto

type ProductVariation struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
}

type ProductVariationRequest struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
}

type StoreProductRequest struct {
	Name              string                    `json:"name"`
	Description       string                    `json:"description"`
	ProductCategoryId uint                      `json:"product_category_id"`
	Variations        []ProductVariationRequest `json:"variations"`
}

type ProductResponse struct {
	ID                uint                    `json:"id"`
	Name              string                  `json:"name"`
	ProductVariations []ProductVariation      `json:"variations"`
	ProductCategory   ProductCategoryResponse `json:"product_category"`
}

type IndexProductResponse struct {
	Products []ProductResponse `json:"products"`
}
