
package utils

import (
	"time"
	"github.com/golang-jwt/jwt"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))


type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID int) string {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func ValidateJWT(tokenStr string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	return claims.UserID, nil
}
