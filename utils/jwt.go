package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func CreateJWT(id int) (string, error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	
	if err != nil {
		log.Println(fmt.Sprintf("Error creating JWT: %s", err.Error()))
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (int, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
			return -1, fmt.Errorf("Invalid Token Signature")
		}
		log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
		return -1, fmt.Errorf("Invalid Token")
	}
	if !tkn.Valid {
		log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
		return -1, fmt.Errorf("Invalid Token")
	}
	return claims.Id, nil
}