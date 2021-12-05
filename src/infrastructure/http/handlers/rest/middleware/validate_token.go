package middleware

import (
	"github.com/labstack/echo/v4"
	"iluvatar/src/shared/utils/jwt"
	"net/http"
)

func ValidateToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			if err := jwt.Token().Validate(token); err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"messsage": err.Error()})
			}
			return next(c)
		}
	}
}
