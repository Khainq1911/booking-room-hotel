package middleware

import (
	"booking-website-be/security"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the token from the cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Println("Token missing in cookie")
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// Verify the token
		token, err := security.VerifyToken(tokenString.Value)
		if err != nil {
			fmt.Printf("Token verification failed: %v\n", err)
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// Store user information in context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		c.Set("user", claims["user"]) // Example: setting a user claim in context

		fmt.Printf("Token verified successfully. Claims: %+v\n", claims)

		
		return next(c)
	}
}
