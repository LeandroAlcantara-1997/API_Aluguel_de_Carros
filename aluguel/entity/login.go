package entity

import (
	"crypto/sha1"
	"fmt"
)

type Admin struct {
	User  string `json:"-"`
	Senha string `json:"senha"`
	Token string `json:"-"`
}

type Login struct {
	Email string `json:"-"`
	Senha string `json:"senha"`
	Token string `json:"-"`
}

func (a *Admin) ValidaAdmin() error {
	if a.User != "admin" {
		return fmt.Errorf("User diferente de admin")
	} else if a.Senha == "" {
		return fmt.Errorf("Senha do admin inválida")
	}
	token, err := GeraToken(a.User + a.Senha)
	a.Token = string(token)
	if err != nil {
		return fmt.Errorf("Erro ao criar token admin %#v", err)
	}

	return nil
}

func GeraToken(dados string) (string, error) {
	sl := sha1.Sum([]byte(dados))
	token := fmt.Sprintf("%x", sl)
	return token, nil
}
