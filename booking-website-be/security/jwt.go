package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const secretKey = "Ftghghttfhgt44"

func GenToken(claims *jwt.MapClaims, ctx echo.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
