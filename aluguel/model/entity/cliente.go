package entity

type Cliente struct {
	Id              int 	`json:"id"`
	Nome            string	`json:"nome"`
	Sobrenome       string	`json:"sobrenome`
	Data_Nascimento string	`json:"data_nascimento"`
	RG              string	`json:"rg`
	CPF             string	`json:"cpf"`
	CNH             string	`json:"cnh"`
	//Contato         Contato 
	/*Endereco        Endereco
	Login			Login*/
}
type Login struct {
	Email	string
	Senha	string
	Token	string
}
type Endereco struct {
	Estado      Estado
	Cidade      string
	Bairro      string
	Logradouro  string
	Rua         string
	Numero      string
	Complemento string
}

type Estado struct {
	ID   int
	Nome string
	Pais string //BRASIL
}

type Contato struct {
	Celular  string
	Telefone string
	Email    string
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
}*/
