package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/LeandroAlcantara-1997/entity"
	"github.com/LeandroAlcantara-1997/repository"
	service "github.com/LeandroAlcantara-1997/service"
	"github.com/gorilla/mux"
)

func AlugarCarro(w http.ResponseWriter, r *http.Request) {
	var aluguel entity.Aluguel
	ids :=mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao ler body ", err)
		return
	}

	aluguel.Id_Cliente, err = strconv.Atoi(ids["cliente"])
	if err != nil {
		service.ReponseError(w, 400, "Erro ao inserir cadastro", err)
	}

	aluguel.Id_Veiculo, err = strconv.Atoi(ids["veiculo"])
	if err != nil {
		service.ReponseError(w, 400, "Erro ao inserir cadastro", err)
	}

	if err = json.Unmarshal(body, &aluguel); err != nil {
		service.ReponseError(w, 400, "Erro ao realizar unmarshal", err)
		return
	}

	veiculo, err := repository.GetVeiculoById(aluguel.Id_Veiculo)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retornar veiculo,", err)
		return
	}

	if err = aluguel.CalculaTotal(veiculo); err != nil {
		service.ReponseError(w, 400, "Erro ao calcular total", err)
		return
	}

	if err = repository.InsertAluguel(aluguel); err != nil {
		service.ReponseError(w, 400, "Erro ao inserir aluguel", err)
		return
	}
	service.JsonResponse(w, 200, aluguel)
}

func GetAlugueis(w http.ResponseWriter, r *http.Request) {
	var alugado []entity.Aluguel
	alugado, err := repository.GetAlugueis()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retonar alugueis", err)
		return
	}
	service.JsonResponse(w, 302, alugado)
}

func CarrosDisponiveis(w http.ResponseWriter, r *http.Request) {
	var veiculos []entity.Veiculo

	veiculos, err := repository.GetCarrosDisponiveis()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retornar veiculos", err)
		return 
	}
	service.JsonResponse(w, 302, veiculos)
}
