package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func AlugarCarro(w http.ResponseWriter, r *http.Request) {
	var aluguel entity.Aluguel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao ler body ", err)
		return
	}

	err = json.Unmarshal(body, &aluguel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao realizar unmarshal", err)
		return
	}
	err = repository.InsertAluguel(aluguel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}
	
	w.WriteHeader(http.StatusOK)
	enconder := json.NewEncoder(w)
	enconder.Encode(aluguel)
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
