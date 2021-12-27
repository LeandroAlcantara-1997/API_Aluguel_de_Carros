package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Redirecionado admin")
}

func GetLoginAdmin(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, "loginAdmin.html", nil)
}

func PostLoginAdmin(w http.ResponseWriter, r *http.Request) {
	var admin entity.Admin
	admin.User = r.FormValue("user")
	admin.Senha = r.FormValue("senha")
	err := admin.ValidaAdmin()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Erro ao validar admin", err)
		return
}

	err = repository.InsertAdmin(&admin)
	if err != nil {
		fmt.Fprint(w, "Erro ao inserir admin: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/homeAdmin", http.StatusFound)

	return
}

func GetByIdCliente(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao converter parametro id para int", err)
		return
	}
	cliente, err := repository.GetByIdCliente(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
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
