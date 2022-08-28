package product

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
)

func IndexProductController(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		products, err := interactor.New(database).IndexProduct(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, products)
	}
}
