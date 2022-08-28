package product

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
	"strconv"
)

func FindProductController(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, _ := strconv.Atoi(c.Param("id"))

		product, err := interactor.New(database).FindProduct(uint(productId))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	}
}
