package router

import (
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/service/authentication"
	"net/http"
	"strings"
)

func apiAuthenticationMiddleware(database db.DB) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			authorizationKey := c.Request().Header.Get("Authorization")

			if authorizationKey == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			tokenData := strings.Split(authorizationKey, " ")

			jwtCredentials, _ := authentication.GetClaimsFromToken(tokenData[1])
			userInfo := jwtCredentials["UserInfo"]

			if userInfo == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			user := userInfo.(map[string]interface{})
			userId, userIdErr := user["user_id"].(float64)
			tokenHash, isCorrect := user["token"].(string)

			if !isCorrect || !userIdErr {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			isValid := database.ValidateTokenExistInStorage(tokenHash, uint(userId))

			if !isValid {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			return next(c)
		}

	}

}

func rolePermissionChecker() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
