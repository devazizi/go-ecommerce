package router

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/delivery/http/v1/auth"
	"go-ecommerce/delivery/http/v1/product"
	"go-ecommerce/delivery/http/v1/product_category"
	"go-ecommerce/delivery/http/v1/shopping_cart"
	"go-ecommerce/validator"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, database db.DB) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found")
	})

	apiv1 := e.Group("api/v1")
	{
		apiV1WithAuth := apiv1.Group("", apiAuthenticationMiddleware(database))
		{
			apiv1.POST("/auth/register", auth.RegisterUserController(database, validator.ValidateRegisterUser))
			apiv1.POST("/auth/login", auth.LoginUserController(database, validator.ValidateLoginUser))
		}

		{
			apiv1.POST("/product-categories", product_category.StoreProductCategoryController(database, validator.ValidateStorePostCategory))
			apiv1.GET("/product-categories", product_category.IndexProductCategoryController(database))
			apiv1.GET("/product-categories/:id", product_category.GetProductCategoryController(database))
		}

		{
			apiv1.GET("/products", product.IndexProductController(database))
			apiv1.GET("/products/:id", product.FindProductController(database))
			apiv1.POST("/products", product.StoreProductController(database, validator.ValidateStoreProduct))
			apiv1.DELETE("/products/:id", product.DestroyProductController(database))
		}

		{
			apiV1WithAuth.GET("/shopping-carts", shopping_cart.GetShoppingCart(database))
		}

		{
			// order
		}

		{
			// payment
		}

	}
}
