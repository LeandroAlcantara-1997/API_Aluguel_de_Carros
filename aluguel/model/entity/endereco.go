package entity

import "fmt"

type Endereco struct {
	id			int
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

func (e *Estado) ValidaEstado() error {
	switch e.Id {
	case 12, 27, 16, 13, 29, 23, 53, 32, 52, 21, 51, 50, 31, 15, 25, 41, 26, 22, 24, 43, 33, 11, 14, 42, 35, 28, 17:
		break
	}
	switch e.Nome {
	case "acre", "alagoas", "amapa", "amazonas", "bahia", "ceara", "distrito federal", "espirito santo", "goias", "maranhao", "mato grosso", "mato grosso do sul", "minas gerais", "para", "paraiba", "pernambuco", "piaui", "rio grande do norte", "rio grande do sul", "rio de janeiro", "rondonia", "roraima", "santa catarina", "sao paulo", "sergipe", "tocantis":
		break
	}

	if e.Pais != "Brasil" {
		return fmt.Errorf("Pais inválido!")
	}
	return nil
}

func (e *Endereco) validaEndereco() error{
	switch e.Logradouro {
		case 
		"aeroporto", "alameda",	"area",	"avenida", "campo",	"chacara", "colonia",	"condominio", "conjunto", "distrito", "esplanada", "estação", "estrada", "favela", "fazenda","feira","jardim","ladeira","lago","lagoa","largo","loteamento","morro","núcleo","parque","passarela","patio","praça","quadra","recanto","residencial","rodovia","rua","setor","sitio","travessa","trecho","trevo","vale","vereda","via","viaduto","viela","vila":
			break
		}
		if e.Bairro == "" {
			return fmt.Errorf("Bairro inválido")
		}else if e.Cidade == "" {
			return fmt.Errorf("Cidade inválida")
		}else if e.Numero == "" {
			return fmt.Errorf("Número inválido")
		}else if e.Rua == "" {
			return fmt.Errorf("Rua inválida")
		}
		return nil
}