package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
)

var redi *redis.Client

type Claims struct {
	Authorized bool
	Id         int64
	Admin      bool
	jwt.StandardClaims
}

func JWTToken(id int64, admin bool) (string, error) {
	claims := &Claims{
		Authorized: true,
		Id:         id,
		Admin:      admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}

	fmt.Println(claims.Admin)

	return token, nil
}

func ValidateToken(token string) error {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("ACCESS_SECRET"), nil 
		})
		fmt.Println("Em baixo do claims")
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return err
			}
			return err
		}

	if !tkn.Valid {
		return fmt.Errorf("Token inválido %v", err)
	}
	return nil
}

func initRedis() {
	redi = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	_, err := redi.Ping().Result()
	if err != nil {
		panic(err)
	}
}
