package entity

import "fmt"

type Endereco struct {
	id          int
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

	if e.Pais != "Brasil" {
		return fmt.Errorf("Pais inválido!")
	}
	return nil
}

func (e *Endereco) validaEndereco() error {
	switch e.Logradouro {
	case
		"aeroporto", "alameda", "area", "avenida", "campo", "chacara", "colonia", "condominio", "conjunto", "distrito", "esplanada", "estação", "estrada", "favela", "fazenda", "feira", "jardim", "ladeira", "lago", "lagoa", "largo", "loteamento", "morro", "núcleo", "parque", "passarela", "patio", "praça", "quadra", "recanto", "residencial", "rodovia", "rua", "setor", "sitio", "travessa", "trecho", "trevo", "vale", "vereda", "via", "viaduto", "viela", "vila":
		break
	}
	if e.Bairro == "" {
		return fmt.Errorf("Bairro inválido")
	} else if e.Cidade == "" {
		return fmt.Errorf("Cidade inválida")
	} else if e.Numero == "" {
		return fmt.Errorf("Número inválido")
	} else if e.Rua == "" {
		return fmt.Errorf("Rua inválida")
	}
	return nil
}

func ValidaEstado(e string) (int, error) {
	switch e {
	case "acre":
		return 12, nil
	case "alagoas":
		return 27, nil
	case "amapa":
		return 16, nil
	case "amazonas":
		return 13, nil
	case "bahia":
		return 29, nil
	case "ceara":
		return 23, nil
	case "distrito federal":
		return 53, nil
	case "espirito santo":
		return 32, nil
	case "goias":
		return 52, nil
	case "maranhao":
		return 21, nil
	case "mato grosso":
		return 51, nil
	case "mato grosso do sul":
		return 50, nil
	case "minas gerais":
		return 31, nil
	case "para":
		return 15, nil
	case "paraiba":
		return 25, nil
	case "parana":
		return 41, nil
	case "pernambuco":
		return 26, nil
	case "piaui":
		return 22, nil
	case "rio grande do norte":
		return 24, nil
	case "rio grande do sul":
		return 43, nil
	case "rio de janeiro":
		return 33, nil
	case "rondonia":
		return 11, nil
	case "roraima":
		return 14, nil
	case "santa catarina":
		return 42, nil
	case "sao paulo":
		return 35, nil
	case "sergipe":
		return 28, nil
	case "tocantis":
		return 17, nil
	}
	return 0, fmt.Errorf("Nome de estado inválido")

}
