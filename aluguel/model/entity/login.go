package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	User  string `json:"user"`
	Senha string	`json:"senha"`
	Token string
}

type Login struct {
	Email string
	Senha string `json:"senha"`
	Token string
}

func (a *Admin) ValidaAdmin() error{
	if a.User != "admin" {
		return fmt.Errorf("User diferente de admin")
	} else if (a.Senha != "admin123456" ) {
		return fmt.Errorf("Senha do admin inválida")
	}
	token, err := GeraToken(a.User + a.Senha)
	a.Token = string(token)
	if err != nil {
		return fmt.Errorf("Erro ao criar token admin ", err)
	}

	return nil
}

func GeraToken(dados string) (string, error){
	senha, err := bcrypt.GenerateFromPassword([]byte(dados), 10)
	if err != nil {
		return "", fmt.Errorf("Erro ao gerar token")
	}

	return string(senha), nil
}