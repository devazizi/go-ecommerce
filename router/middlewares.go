package router

import "github.com/labstack/echo/v4"

func apiAuthenticationMiddleware() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func rolePermissionChecker() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
