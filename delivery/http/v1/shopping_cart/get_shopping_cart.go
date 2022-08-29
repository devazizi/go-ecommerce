package shopping_cart

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
)

func GetShoppingCart(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		shoppingCart, err := interactor.New(database).GetShoppingCart(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, shoppingCart)
	}
}
