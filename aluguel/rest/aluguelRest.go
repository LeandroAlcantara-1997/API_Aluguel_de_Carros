package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	service "github.com/LeandroAlcantara-1997/service"
	"github.com/LeandroAlcantara-1997/entity"
	"github.com/LeandroAlcantara-1997/repository"
)

func AlugarCarro(w http.ResponseWriter, r *http.Request) {
	var aluguel entity.Aluguel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao ler body ", err)
		return
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
