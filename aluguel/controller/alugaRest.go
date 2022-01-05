package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func AlugarCarro(w http.ResponseWriter, r *http.Request) {
	
	return
}

func GetAlugueis(w http.ResponseWriter, r *http.Request) {
	var alugado []entity.Aluguel
	alugado, err := repository.GetAlugueis()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(alugado)
	return
}
