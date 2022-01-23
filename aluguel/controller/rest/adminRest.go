package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	service "github.com/LeandroAlcantara-1997/controller/service"
	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	GetClientesCadastrados(w, r)
	fmt.Fprint(w, "Redirecionado admin")
}

func GetLoginAdmin(w http.ResponseWriter, r *http.Request) {
	service.ExecuteTemplate(w, "loginAdmin.html", nil)
}

func PostLoginAdmin(w http.ResponseWriter, r *http.Request) {
	var admin entity.Admin
	admin.User = r.FormValue("user")
	admin.Senha = r.FormValue("senha")
	err := admin.ValidaAdmin()
	if err != nil {
		service.ReponseError(w, 401,  "Erro ao validar admin", err)
		return
	}

	err = repository.LogarAdmin(&admin)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao validar admin", err )
		return
	}

	http.Redirect(w, r, "/admin/home", http.StatusFound)
}

func GetByIdCliente(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao converter parametro id para int", err )
		return
	}
	cliente, err := repository.GetByIdCliente(id)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao consultar cliente", err)
		return
	}
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(cliente)
}

func GetClientesCadastrados(w http.ResponseWriter, r *http.Request) {
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
	modelo := r.FormValue("modelo")
	marca := r.FormValue("marca")
	ano := r.FormValue("ano")
	cor := r.FormValue("cor")
	km_Litro, err := strconv.ParseFloat(r.FormValue("km_litro"), 64)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao converter km_litro para float", err)
		return
	}

	valor_Dia, err := strconv.ParseFloat(r.FormValue("valor_dia"), 64)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao converter valor_dia para float", err)
		return
	}
	valor_Hora, err := strconv.ParseFloat(r.FormValue("valor_hora"), 64)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao converter valor_hora para float", err)
		return
	}

	veiculo := entity.Veiculo{
		Modelo:     modelo,
		Marca:      marca,
		Ano:        ano,
		Cor:        cor,
		Km_Litro:   km_Litro,
		Valor_Dia:  valor_Dia,
		Valor_Hora: valor_Hora,
	}

	if err = veiculo.ValidaVeiculo(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	
	if err = repository.InsertVeiculo(&veiculo); err != nil {
		service.ReponseError(w, 400, "Erro ao inserir veiculo", err )
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Carro cadastrado com sucesso!", veiculo)
}

func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request) {
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao retonar carros", err)
		return
	}

	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculos)
}
