package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeandroAlcantara-1997/model/repository"
)

func GetByIdCliente(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("id")
	id, err := strconv.Atoi(value)
	fmt.Fprint(w, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao converter parametro id para int", err)
		return
	}
	cliente, err := repository.GetByIdCliente(1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao retornar cadastro", err)
		return
	}
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(cliente)
	return
}

func GetClientesCadastrados(w http.ResponseWriter, r *http.Request) {
	clientes, err := repository.GetClientesCadastrados()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(clientes)
	return
}
func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request) {
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculos)
	return
}
