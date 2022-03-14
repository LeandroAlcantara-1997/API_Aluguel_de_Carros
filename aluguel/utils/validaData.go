package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ValidaData(date string) (time.Time, error) {
	if len(date) != 10 {
		return time.Time{},  fmt.Errorf("Data com formato incorreto")
	}
	
	dataArray := strings.Split(date, "/")
	dia, err := strconv.Atoi(dataArray[0])
	if err != nil {
		return time.Time{} , fmt.Errorf("Erro ao parsear ano para int %v", err)
	}

	m, err := strconv.Atoi(dataArray[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao parsear mes para int %v", err)
	}

	mes, err := ValidaMes(m)
	if err != nil {
		return time.Time{}, fmt.Errorf("Error: %v", err) 
	}

	ano, err := strconv.Atoi(dataArray[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("Erro ao parsear ano para int %v", err)

	}

	var data = time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)

	return data, nil
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