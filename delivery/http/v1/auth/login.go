package auth

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/contract"
	"go-ecommerce/dto"
	"go-ecommerce/interactor"
	"net/http"
)

func LoginUserController(database db.DB, validator contract.ValidateLoginUser) echo.HandlerFunc {
	return func(e echo.Context) error {
		requestStruct := dto.LoginUserRequest{}

		if reqErr := e.Bind(&requestStruct); reqErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		if validateErr := validator(requestStruct); validateErr != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, validateErr.Error())
		}

		_, err := interactor.New(database).LoginUser(e.Request().Context(), requestStruct)
		if err != nil {
			return err
		}

		return nil
	}
}
