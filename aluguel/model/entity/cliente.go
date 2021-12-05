package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type Cliente struct {
	Id              int      `json:"id"`
	Nome            string   `json:"nome"`
	Sobrenome       string   `json:"sobrenome"`
	Data_Nascimento string   `json:"data_nascimento"`
	RG              string   `json:"rg"`
	CPF             string   `json:"cpf"`
	CNH             string   `json:"cnh"`
	Contato         Contato  `json:"contato"`
	Endereco        Endereco `json:"endereco"`
	Login           Login    `json:"login"`
}

type Contato struct {
	id       int
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
	} else if c.Login.Senha == "" {
		return fmt.Errorf("A senha não pode estar vazia")
	} else if c.CPF == "" || len(c.CPF) != 10 {
		return fmt.Errorf("Número de CPF inválido")
	}

	data, err := ValidaData(c.Data_Nascimento)
	if err != nil {
		return fmt.Errorf("%v ", err)
	}
	c.Data_Nascimento = data
	err = c.Contato.ValidaContato()
	if err != nil {
		return fmt.Errorf(" %v", err)
	}
	c.Login.Email = c.Contato.Email
	passbyte, err := bcrypt.GenerateFromPassword([]byte(c.Login.Email+c.Login.Senha), 10)

	if err != nil {
		return fmt.Errorf("Erro ao cadastrar token")
	}
	c.Login.Token = string(passbyte)

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

func validaNumber(number string) bool {
	for _, value := range number {
		if !unicode.IsDigit(value) {
			return true
		}
	}
	return false
}

func ValidaData(date string) (string, error) {
	data := strings.Split(date, "/")
	dia, err := strconv.Atoi(data[0])
	if err != nil {
		return "", fmt.Errorf("Erro ao parsear ano para int %v", err)
	}

	m, err := strconv.Atoi(data[1])
	if err != nil {
		return "", fmt.Errorf("Erro ao parsear mes para int %v", err)
	}

	mes, err := ValidaMes(m)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	ano, err := strconv.Atoi(data[2])
	if err != nil {
		return "", fmt.Errorf("Erro ao parsear ano para int %v", err)

	}

	var nasc = time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)
	compare := time.Now()
	if !nasc.Before(compare) {
		return "", fmt.Errorf("Ano de nascimento deve ser anterior a data atual")
	}

	return nasc.Format("2006-01-02"), nil
}

func ValidaMes(mes int) (time.Month, error) {
	switch mes {
	case 1:
		return 1, nil
	case 2:
		return 2, nil
	case 3:
		return 3, nil
	case 4:
		return 4, nil
	case 5:
		return 5, nil
	case 6:
		return 6, nil
	case 7:
		return 7, nil
	case 8:
		return 8, nil
	case 9:
		return 9, nil
	case 10:
		return 10, nil
	case 11:
		return 11, nil
	case 12:
		return 12, nil
	default:
		return 0, fmt.Errorf("Indice de mes inválido")

	}
}

/*{
    "id": 1,
	"nome": "Leandro",
	"sobrenome": "Alcantara",
	"data_nascimento": "24/12/1991",
	"rg": "4462546215",
	"cpf": "456891621651",
	"cnh": "15181516516",
	"contato": {
		"celular": "5151515",
		"telefone": "518162480",
		"email": "21484848"
	},
	"endereco": {
		"estado": {
			"nome": "são paulo",
			"pais": "Brasil"
		},
		"cidade": "Taboão da Serra",
		"bairro": "Jd.Elisabete",
		"logradouro": "rua",
		"rua": "Almeida filho",
		"numero": "52",
		"complemento": "casa 06"
	},
	"login": {
		"senha": "123456"
	}
}*/
