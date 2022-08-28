package db

func (db DB) IndexProducts() ([]Product, error) {
	var products []Product

	db.store.Preload("ProductVariations").Preload("ProductCategory").Find(&products)

	return products, nil
}
