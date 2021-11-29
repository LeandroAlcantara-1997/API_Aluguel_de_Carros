package entity


type Cliente struct {
	Id              int 	`json:"id"`
	Nome            string	`json:"nome"`
	Sobrenome       string	`json:"sobrenome`
	Data_Nascimento string	`json:"data_nascimento"`
	RG              string	`json:"rg`
	CPF             string	`json:"cpf"`
	CNH             string	`json:"cnh"`
	//Contato         Contato `json:"contato"`
	/*Endereco        Endereco `json:"endereco"`
	Login			Login	   `json:"login"`*/
}
type Login struct {
	Email	string	`json:"email`
	Senha	string	`json:"senha"`
	token	string
}
type Endereco struct {
	Estado      Estado	`json:"estado"`
	Cidade      string	`json:"cidade"`
	Bairro      string	`json:"bairro"`
	Logradouro  string	`json:"logradouro"`
	Rua         string	`json:"rua"`
	Numero      string	`json:"numero"`
	Complemento string	`json:"complemento"`
}

type Estado struct {
	ID   int	
	Nome string	`json:"nome"`
	Pais string `json:"pais"`//BRASIL
}

type Contato struct {
	Celular  string `json:"celular"`
	Telefone string	`json:"telefone"`
	Email    string	`json:"email"`
}


/*var c Cliente {
	Id: 1,
	Nome: "Leandro",
	Sobrenome: "Alcantara",
	Data_Nascimento: "24/12/1991",
	RG: "4462546215",
	CPF: "456891621651",
	CNH: "15181516516",
	Contato: "11 97261451",
}


{
    "id": 1,
	"nome": "Leandro",
	"sobrenome": "Alcantara",
	"data_nascimento": "24/12/1991",
	"rg": "4462546215",
	"cpf": "456891621651",
	"cnh": "15181516516"
   }*/
