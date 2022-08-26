package router

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/delivery/http/v1/auth"
	"go-ecommerce/validator"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, database db.DB) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found")
	})

	apiv1 := e.Group("api/v1")
	{
		{
			apiv1.POST("/auth/register", auth.RegisterUserController(database, validator.ValidateRegisterUser))
			//apiV1.POST("/auth/login")
		}

		{
			// product

		}

		{
			// shopping cart
		}

		{
			// order
		}

		{
			// payment
		}

	}
}
