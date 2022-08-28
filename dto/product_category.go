package dto

type StorePostCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductCategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IndexProductCategoriesResponse struct {
	ProductCategories []ProductCategoryResponse `json:"product_categories"`
}
