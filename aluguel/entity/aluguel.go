package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	utils "github.com/LeandroAlcantara-1997/utils"
)

type Aluguel struct {
	Id_Cliente  int  `json:"cliente,omitempty"`
	Id_Veiculo  int  `json:"veiculo,omitempty"`
	Inicio      string `json:"inicio"`
	Retorno     string `json:"retorno"`
	Valor_Total float64
}

func (aluguel *Aluguel) CalculaTotal(veiculo Veiculo) error {
	inicio, err := StringToTime(aluguel.Inicio)
	if err != nil {
		return err
	}
	retorno, err := StringToTime(aluguel.Retorno)
	if err != nil {
		return err
	}
	totalDia := retorno.Day() - inicio.Day()
	total := veiculo.Valor_Dia * float64(totalDia)
	aluguel.Valor_Total = total

	return nil
}

func StringToTime(date string) (time.Time, error) {
	dataFormat := strings.Split(date, "/")
	dataHora := strings.Split(dataFormat[2], " ")
	horaFormat := strings.Split(dataHora[1], ":")

	ano, err := strconv.Atoi(dataFormat[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter ano: %#v", err)
	}

	mes, err := strconv.Atoi(dataFormat[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter mes: %#v", err)
	}

	dia, err := strconv.Atoi(dataHora[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter dia: %#v", err)
	}

	hora, err := strconv.Atoi(horaFormat[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter hora: %#v", err)
	}

	minuto, err := strconv.Atoi(horaFormat[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter minuto: %#v", err)
	}

	segundo, err := strconv.Atoi(horaFormat[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter segundo: %#v", err)
	}

	mesTime, err := utils.ValidaMes(mes)
	if err != nil {
		return time.Time{}, err
	}

	tempoInicio := time.Date(ano, mesTime, dia, hora, minuto, segundo, 0, time.UTC)

	return tempoInicio, nil
}

func (aluguel *Aluguel) verificaData() error {
	inicio, err := utils.ValidaData(aluguel.Inicio)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}

	retorno, err := utils.ValidaData(aluguel.Retorno)
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

	if !inicio.Before(retorno) {
		return fmt.Errorf("Data do retorno do aluguel deve ser futura a data inicio")
	}
	return nil
}
