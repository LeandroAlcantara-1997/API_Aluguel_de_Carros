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

	return token, nil
}

func ValidaCliente(token string) error{
	claim, err := validateToken(token)
	if err != nil {
		return err
	}

	if claim.Admin {
		return fmt.Errorf("Só clientes podem acessar essa rota")
	}

	return nil
}

func ValidaAdmin(token string) error {
	claim, err := validateToken(token)
	if err != nil {
		return err
	}

	if !claim.Admin {
		return fmt.Errorf("Só admins podem acessar essa rota")
	}

	return nil
}

func validateToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("ACCESS_SECRET"), nil 
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return claims, err
			}
			return claims, err
		}

	if !tkn.Valid {
		return claims, fmt.Errorf("Token inválido %v", err)
	}
	return claims, nil
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
