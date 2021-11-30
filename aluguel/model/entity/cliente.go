package entity

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type Cliente struct {
	Id              int     `json:"id"`
	Nome            string  `json:"nome"`
	Sobrenome       string  `json:"sobrenome"`
	Data_Nascimento string  `json:"data_nascimento"`
	RG              string  `json:"rg"`
	CPF             string  `json:"cpf"`
	CNH             string  `json:"cnh"`
	Contato         Contato `json:"contato"`
	//Endereco        Endereco `json:"endereco"`
	Login Login `json:"login"`
}
type Login struct {
	email string
	Senha string `json:"senha"`
	token string
}
type Endereco struct {
	Estado      Estado `json:"estado"`
	Cidade      string `json:"cidade"`
	Bairro      string `json:"bairro"`
	Logradouro  string `json:"logradouro"`
	Rua         string `json:"rua"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
}

type Estado struct {
	Id   int
	Nome string `json:"nome"`
	Pais string `json:"pais"` //BRASIL
}

type Contato struct {
	Id		 int	`json:id_contato`
	Celular  string `json:"celular"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

func (c *Cliente) ValidaCliente() error {
	if c.Nome == "" || validaText(c.Nome){
		return fmt.Errorf("Nome inválido!")
	} else if c.Sobrenome == "" || validaText(c.Sobrenome) {
		return fmt.Errorf("Sobrenome inválido")
	} else if c.CNH == "" || validaNumber(c.CNH) {
		return fmt.Errorf("Número de CNH inválida")
	} else if c.Data_Nascimento == "" {
		return fmt.Errorf("Data de nascimento inválida")
	} else if c.RG == "" || validaNumber(c.RG){
		return fmt.Errorf("Número de RG inválido")
	} else if c.Login.Senha == "" {
		return fmt.Errorf("A senha não pode estar vazia")
	}
	c.Contato.Id = c.Id
	err := c.Contato.ValidaContato()
	if err != nil {
		return fmt.Errorf(" %v", err)
	}
	c.Login.email = c.Contato.Email
	passbyte, err := bcrypt.GenerateFromPassword([]byte(c.Login.email+c.Login.Senha), 10)
	fmt.Println(string(passbyte))

	if err != nil {
		return fmt.Errorf("Erro ao cadastrar token")
	}
	c.Login.token = string(passbyte)

	return nil
}

func (c *Contato) ValidaContato() error {
	if c.Celular == "" || validaNumber(c.Celular){
		return fmt.Errorf("É necessário informar o número de celular")
	} else if c.Email == "" {
		return fmt.Errorf("Email inválido")
	} else if c.Telefone == "" || validaNumber(c.Telefone) {
		return fmt.Errorf("Número de telefone inválido")
	}
	return nil
}

func (e *Estado) ValidaEstado() error {
	switch e.Id {
	case 12, 27, 16, 13, 29, 23, 53, 32, 52, 21, 51, 50, 31, 15, 25, 41, 26, 22, 24, 43, 33, 11, 14, 42, 35, 28, 17:
		break
	}
	switch e.Nome {
	case "acre", "alagoas", "amapa", "amazonas", "bahia", "ceara", "distrito federal", "espirito santo", "goias", "maranhao", "mato grosso", "mato grosso do sul", "minas gerais", "para", "paraiba", "pernambuco", "piaui", "rio grande do norte", "rio grande do sul", "rio de janeiro", "rondonia", "roraima", "santa catarina", "sao paulo", "sergipe", "tocantis":
		break
	}
	return nil
}
func validaText(text string) bool{
	for _, value := range text{
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
	"login": {
		"senha": "123456"
	}
}*/
