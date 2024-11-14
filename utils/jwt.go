package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
)

// Custom Claims structure to include user-specific data
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for the given user ID
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(config.GetJWTExpirationTime())
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Signing the token with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWTSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

// ValidateJWT validates the token and returns the claims if valid
func ValidateJWT(tokenStr string) (*Claims, error) {
	// Parse the token with the claims
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key for verification
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token: %v", err)
	}

	// Type assertion to extract claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	return claims, nil
}
