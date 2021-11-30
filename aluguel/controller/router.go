package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func CadastraCliente(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body) 
	if err != nil {
		log.Fatal("Erro ao cadastrar Cliente ", err)
	}
	var novocadastro entity.Cliente
	err = json.Unmarshal(body, &novocadastro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao converter json para cliente: ", err)
		return
	}
	err = novocadastro.ValidaCliente()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao cadastrar cliente: ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Cadastrado realidado com sucesso!")
}

func LoginCliente(w http.ResponseWriter, r *http.Request){
	return 
}

func RestauraSenha(w http.ResponseWriter, r *http.Request){
	return 
}

func AlugarCarro(w http.ResponseWriter, r *http.Request){
	return 
}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request){
	return 
}

func CadastroCarros(w http.ResponseWriter, r *http.Request){
	return 
}

func GetCarrosCadastrados(w http.ResponseWriter, r *http.Request){
	return 
}


func GetAluguel(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		c := entity.Veiculo {
			Id: 1,
			Modelo: "Corsa",
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(c)
	} else {
		return 
	}
}


