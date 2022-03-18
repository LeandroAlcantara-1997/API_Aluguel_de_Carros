package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

func ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", t.Header["alg"])
		}
		return []byte("ACCESS_SECRET"), nil
	})

	return err == nil
}

func JWTToken(id int64, admin bool) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["admin"] = admin
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}

	fmt.Println(atClaims["admin"])

	return token, nil
}