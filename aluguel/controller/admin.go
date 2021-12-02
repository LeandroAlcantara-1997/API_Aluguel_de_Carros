package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeandroAlcantara-1997/model/repository"
	"github.com/gorilla/mux"
)


func GetByIdCliente(w http.ResponseWriter, r *http.Request){
	value := mux.Vars(r) 
	id, err := strconv.Atoi(value["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w ,"Erro ao converter parametro id para int", err)
		return 
	}
	row, err := repository.GetByIdCliente(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao retornar cadastro", err)
		return 
	}
	w.WriteHeader(http.StatusFound)
	fmt.Fprint(w, row)
	return
}

func CadastroCarros(w http.ResponseWriter, r *http.Request) {
	return
}

func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request){
	return 
}