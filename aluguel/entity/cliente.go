package entity

import (
	"fmt"
	"time"
	"unicode"


	"github.com/LeandroAlcantara-1997/utils"
)

type Cliente struct {
	Id              int64   `json:"id"`
	Nome            string   `json:"nome"`
	Sobrenome       string   `json:"sobrenome"`
	Data_Nascimento string   `json:"data_nascimento"`
	RG              string   `json:"rg"`
	CPF             string   `json:"cpf"`
	CNH             string   `json:"cnh"`
	Contato         Contato  `json:"contato"`
	Endereco        Endereco `json:"endereco"`
	Login           Login    `json:"login,omitempty"`
}

type Contato struct {
	Id       int	`json:"-"`
	Celular  string `json:"celular"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

func (c *Cliente) ValidaCliente() error {
	if c.Nome == "" || validaText(c.Nome) {
		return fmt.Errorf("Nome inválido!")
	} else if c.Sobrenome == "" || validaText(c.Sobrenome) {
		return fmt.Errorf("Sobrenome inválido")
	} else if c.CNH == "" || validaNumber(c.CNH) {
		return fmt.Errorf("Número de CNH inválida")
	} else if c.Data_Nascimento == "" {
		return fmt.Errorf("Data de nascimento inválida")
	} else if c.RG == "" || validaNumber(c.RG) {
		return fmt.Errorf("Número de RG inválido")
	} else if c.CPF == "" || len(c.CPF) != 10 {
		return fmt.Errorf("Número de CPF inválido")
	}
	if err := c.Login.validaSenha(); err != nil {
		return err
	}
	data, err := ValidaDataNascimento(c.Data_Nascimento)
	if err != nil {
		return fmt.Errorf("%v ", err)
	}
	c.Data_Nascimento = data
	err = c.Contato.ValidaContato()
	if err != nil {
		return fmt.Errorf(" %v", err)
	}
	c.Login.Email = c.Contato.Email

	c.Login.Token, err = GeraToken(c.Login.Senha)
	if err != nil {
		return fmt.Errorf("Erro ao cadastrar token")
	}
	
	return nil
}

func (c *Contato) ValidaContato() error {
	if c.Celular == "" || validaNumber(c.Celular) {
		return fmt.Errorf("É necessário informar o número de celular")
	} else if c.Email == "" {
		return fmt.Errorf("Email inválido")
	} else if c.Telefone == "" || validaNumber(c.Telefone) {
		return fmt.Errorf("Número de telefone inválido")
	}
	return nil
}

func validaText(text string) bool {
	for _, value := range text {
		if unicode.IsDigit(value) {
			return true
		}
	}
	return false
}

func ValidaDataNascimento(date string) (string, error){
	nasc, err := utils.ValidaData(date)
	if err != nil {
		return "", fmt.Errorf("%#v", err)
	}

	compare := time.Now()
	if !nasc.Before(compare) {
		return "", fmt.Errorf("Ano de nascimento deve ser anterior a data atual")
	}

	return nasc.Format("2006-01-02"), nil
}

func validaNumber(number string) bool {
	for _, value := range number {
		if !unicode.IsDigit(value) {
			return true
		}
	}
	return false
}
