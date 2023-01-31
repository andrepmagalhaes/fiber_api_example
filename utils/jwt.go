package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id int `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func CreateJWT(id int, userType string) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Id: id,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	
	if err != nil {
		log.Printf("Error creating JWT: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(c *fiber.Ctx) (int, string, error) {

	secret := os.Getenv("JWT_SECRET")

	authorization := c.Get("Authorization")
	if authorization == "" {
		return -1, "", fmt.Errorf("Unauthorized")
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("Error verifying JWT: %s", err.Error())
			return -1, "", fmt.Errorf("invalid token signature")
		}
		log.Printf("Error verifying JWT: %s", err.Error())
		return -1, "", fmt.Errorf("invalid token")
	}
	if !tkn.Valid {
		log.Printf("Error verifying JWT: %s", err.Error())
		return -1, "", fmt.Errorf("invalid token")
	}
	return claims.Id, claims.UserType, nil
}