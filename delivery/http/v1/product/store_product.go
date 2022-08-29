package product

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/contract"
	"go-ecommerce/dto"
	"go-ecommerce/interactor"
	"net/http"
)

func StoreProductController(database db.DB, validator contract.ValidateStoreProduct) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := dto.StoreProductRequest{}
		if reqErr := c.Bind(&request); reqErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		if validationErr := validator(request); validationErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, validationErr.Error())
		}

		product, err := interactor.New(database).StoreProduct(c.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, product)
	}
}
