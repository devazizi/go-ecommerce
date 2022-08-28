package product_category

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/contract"
	"go-ecommerce/dto"
	"go-ecommerce/interactor"
	"net/http"
)

func StoreProductCategoryController(database db.DB, validator contract.ValidateStorePostCategory) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := dto.StorePostCategoryRequest{}
		if reqErr := c.Bind(&request); reqErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		if validationErr := validator(request); validationErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, validationErr.Error())
		}

		productCategory, err := interactor.New(database).StoreProductCategory(c.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, productCategory)
	}
}
