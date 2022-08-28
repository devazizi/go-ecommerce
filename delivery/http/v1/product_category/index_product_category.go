package product_category

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/interactor"
	"net/http"
)

func IndexProductCategoryController(database db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		productCategories, err := interactor.New(database).IndexProductCategory(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, productCategories)
	}
}
