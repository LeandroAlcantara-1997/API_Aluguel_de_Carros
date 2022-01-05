package entity

import (
	"fmt"
	"time"

	utils "github.com/LeandroAlcantara-1997/controller/utils"
)

type Aluguel struct {
	Id_Cliente   int    `json:"cliente`
	Id_Veiculo   int    `json:"veiculo"`
	Data_Inicio  string `json:"inicio`
	Data_Retorno string `json:"retorno"`
	Valor_Total  float64
}

func CalculaTotal(veiculo Veiculo) (float64, error) {
	return 0, nil
}

func (aluguel *Aluguel) verificaData() error {
	inicio, err := utils.ValidaData(aluguel.Data_Inicio)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}

	retorno, err := utils.ValidaData(aluguel.Data_Retorno)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}

	atual := time.Now()
	if inicio.Before(atual) {
		return fmt.Errorf("Data do aluguel deve ser futura a data atual")
	}

	if retorno.Before(atual) {
		return fmt.Errorf("Data do retorno do aluguel deve ser futura a data atual")
	}

	if 	!inicio.Before(retorno) {
		return fmt.Errorf("Data do retorno do aluguel deve ser futura a data inicio")
	}
	return nil
}
