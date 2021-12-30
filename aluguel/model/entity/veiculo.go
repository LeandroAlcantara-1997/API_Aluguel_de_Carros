package entity

import (
	"fmt"
	"unicode"
)

type Veiculo struct {
	Id         int64   `json:"omitempty"`
	Modelo     string  `json:"modelo"`
	Marca      string  `json:"marca"`
	Ano        string  `json:"ano"`
	Cor        string  `json:"cor"`
	Km_Litro   float64 `json:"km_litro"`
	Valor_Dia  float64 `json: "valor_dia"`
	Valor_Hora float64 `json: "valor_hora"`
}

func (veiculo *Veiculo) ValidaVeiculo() error {
	if veiculo.Modelo == "" {
		return fmt.Errorf("Campo modelo não pode estar vazio")
	} else if veiculo.Marca == "" {
		return fmt.Errorf("Campo marca não pode estar vazio")
	}else if veiculo.Km_Litro == 0 {
		return fmt.Errorf("A quantidade de km por litro não pode ser 0")
	}else if veiculo.Valor_Dia == 0 {
		return fmt.Errorf("O valor dia não pode ser igual a 0")
	}else if veiculo.Valor_Hora == 0 {
		return fmt.Errorf("O valor hora não pode ser igual a 0")
	}
	if err := validaAno(veiculo.Ano); err != nil {
		return fmt.Errorf("", err)
	}
	return nil
}

func validaAno(ano string) error {

	if ano == "" {
		return fmt.Errorf("O campo ano não pode estar vazio")
	}
	if len(ano) != 4 {
		return fmt.Errorf("O campo ano deve conter 4 digitos")
	} 
	for _, value := range ano {
		if !unicode.IsDigit(value) {
			return fmt.Errorf("O campo ano só pode conter digitos")
		}
	}
	return nil
}
