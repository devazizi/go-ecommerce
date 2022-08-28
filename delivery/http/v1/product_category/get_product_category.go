package product_category

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
	"strconv"
)

func GetProductCategoryController(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		productCategoryId, _ := strconv.Atoi(c.Param("id"))

		productCategories, err := interactor.New(database).FindProductCategory(c.Request().Context(), uint(productCategoryId))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, productCategories)
	}
}
