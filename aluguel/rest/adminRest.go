package rest

import (
	"encoding/json"
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
	id := mux.Vars(r)

	cliente, err := repository.GetClienteById(id["id"])
	if err != nil {
		service.ReponseError(w, 400, "Erro ao consultar cliente", err)
		return
	}
	service.JsonResponse(w, 302, cliente)
}

func GetClientesCadastrados(w http.ResponseWriter, r *http.Request) {
	clientes, err := repository.GetClientesCadastrados()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retornar cliente", err)
		return
	}
	service.JsonResponse(w, 302, clientes)
}

func CadastraCarro(w http.ResponseWriter, r *http.Request) {
	var veiculo entity.Veiculo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao ler body", err)
		return
	}

	if err = json.Unmarshal(body, &veiculo); err != nil {
		service.ReponseError(w, 400, "Erro ao realizar unmarshal", err)
		return
	}

	if err = veiculo.ValidaVeiculo(); err != nil {
		service.ReponseError(w, 400, "Erro ao realizar unmarshal", err)
		return
	}

	if err = repository.InsertVeiculo(&veiculo); err != nil {
		service.ReponseError(w, 400, "Erro ao inserir veiculo", err)
		return
	}
	service.JsonResponse(w, 201, veiculo)
}

func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request) {
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retonar carros", err)
		return
	}
	service.JsonResponse(w, 302, veiculos)
}
