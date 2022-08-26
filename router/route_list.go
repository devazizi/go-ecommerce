package router

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/delivery/http/v1/auth"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found")
	})

	apiv1 := e.Group("api/v1")
	{
		{
			// auth
			apiv1.POST("/auth/register", auth.RegisterUserController())
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
