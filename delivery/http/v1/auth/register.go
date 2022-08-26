package auth

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/contract"
	"go-ecommerce/dto"
	"go-ecommerce/interactor"
	"net/http"
)

func RegisterUserController(database db.DB, validator contract.ValidateRegisterUser) echo.HandlerFunc {

	return func(e echo.Context) error {

		req := dto.RegisterUserRequest{}

		if err := e.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		userResponse, err := interactor.New(database).RegisterUser(e.Request().Context(), req)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "some err happen")
		}

		return e.JSON(http.StatusCreated, userResponse)
	}
}
