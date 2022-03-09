package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LeandroAlcantara-1997/entity"
	"github.com/LeandroAlcantara-1997/repository"
	service "github.com/LeandroAlcantara-1997/service"
	"github.com/gorilla/mux"
)

func PostLoginAdmin(w http.ResponseWriter, r *http.Request) {
	var admin entity.Admin
	admin.User = r.FormValue("user")
	admin.Senha = r.FormValue("senha")

	if err := admin.ValidaAdmin(); err != nil {
		service.ReponseError(w, 401, "Erro ao validar admin", err)
		return
	}

	if err := repository.LogarAdmin(&admin); err != nil {
		service.ReponseError(w, 400, "Erro ao validar admin", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func GetClienteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	cliente, err := repository.GetByIdCliente(id["id"])
	if err != nil {
		service.ReponseError(w, 400, "Erro ao consultar cliente", err)
		return
	}
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(cliente)
}

func GetClientesCadastrados(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientes, err := repository.GetClientesCadastrados()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retornar cliente", err)
		return
	}

	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(clientes)
}

func CadastraCarro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var veiculo entity.Veiculo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler body"))
		return
	}
	
	if err = json.Unmarshal(body, &veiculo); err != nil {
		w.Write([]byte("Erro ao realizar unmarshal"))
		return
	}

	if err = veiculo.ValidaVeiculo(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if err = repository.InsertVeiculo(&veiculo); err != nil {
		service.ReponseError(w, 400, "Erro ao inserir veiculo", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculo)
}

func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retonar carros", err)
		return
	}

	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculos)
}
