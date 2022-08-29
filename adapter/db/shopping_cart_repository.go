package db

import "gorm.io/gorm/clause"

func (db DB) GetShoppingCart(userId uint) ShoppingCart {
	var shoppingCart ShoppingCart
	db.store.
		Preload("ShoppingCartItems.ProductVariation.Product.ProductCategory").
		Preload(clause.Associations).
		FirstOrCreate(&shoppingCart, ShoppingCart{UserID: userId})

	return shoppingCart
}
