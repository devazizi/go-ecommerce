package product

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
	"strconv"
)

func DestroyProductController(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		productId, _ := strconv.Atoi(c.Param("id"))

		err := interactor.New(database).DestroyProduct(c.Request().Context(), uint(productId))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, "success")
	}
}
