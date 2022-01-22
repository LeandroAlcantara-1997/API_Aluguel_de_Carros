package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	service "github.com/LeandroAlcantara-1997/controller/service"
)

type Aluguel struct {
	Id_Cliente  int64  `json:"cliente"`
	Id_Veiculo  int64  `json:"veiculo"`
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
	fmt.Println(totalDia)
	total := veiculo.Valor_Dia * float64(totalDia)
	fmt.Println(total)
	aluguel.Valor_Total = total
	fmt.Println(total)

	return nil
}

func StringToTime(date string) (time.Time, error) {
	dataFormat := strings.Split(date, "/")
	dataHora := strings.Split(dataFormat[2], " ")
	horaFormat := strings.Split(dataHora[1], ":")

	ano, err := strconv.Atoi(dataFormat[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter ano: ", err)
	}

	mes, err := strconv.Atoi(dataFormat[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter mes: ", err)
	}

	dia, err := strconv.Atoi(dataHora[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter dia: ", err)
	}

	hora, err := strconv.Atoi(horaFormat[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter hora: ", err)
	}

	minuto, err := strconv.Atoi(horaFormat[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter minuto: ", err)
	}

	segundo, err := strconv.Atoi(horaFormat[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao converter segundo: ", err)
	}

	mesTime, err := service.ValidaMes(mes)
	if err != nil {
		return time.Time{}, err
	}

	tempoInicio := time.Date(ano, mesTime, dia, hora, minuto, segundo, 0, time.UTC)

	return tempoInicio, nil
}

func (aluguel *Aluguel) verificaData() error {
	inicio, err := service.ValidaData(aluguel.Inicio)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}

	retorno, err := service.ValidaData(aluguel.Retorno)
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
