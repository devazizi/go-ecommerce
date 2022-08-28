package dto

type ProductVariation struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
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
