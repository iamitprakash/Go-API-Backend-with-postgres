package services

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username  string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
	Role      string `json:"roles"`
}

// Valid implements jwt.Claims.
func (Claims) Valid() error {
	panic("unimplemented")
}

func GenerateBearerToken(claims Claims) (string, error) {
	// Create a new JWT signing method using the HS256 algorithm.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign the token with the secret key.
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", tokenString), nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	// Parse the token.
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key.
		return []byte("my-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid.
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Get the claims from the token.
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}
