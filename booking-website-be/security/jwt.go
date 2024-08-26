package security

import (
	"booking-website-be/model"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("Ftghghttfhgt44")

func GenerateJWTToken(user model.SignIn) (string, error) {
	// Define token claims including the user's role
	claims := jwt.MapClaims{
		"user_id": user.User_id,
		"phone":   user.Phone,
		"role":    user.Role,
		// additional claims
	}
	// Generate and sign the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(tokenString)
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
