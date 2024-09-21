package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken := c.Request().Header.Get("authorization")
		token := strings.Split(reqToken, " ")[1]
		fmt.Println(token)
		claims := jwt.MapClaims{}
		fmt.Println(claims)
		_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte("Ftghghttfhgt44"), nil
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "unable to parse token")
		}

		if !claims["is_admin"].(bool) {
			return c.JSON(http.StatusInternalServerError, "unable to access")
		}
		return next(c)
	}
}
