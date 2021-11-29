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
	if r.Method == "POST"{
		w.WriteHeader(http.StatusCreated)
		body, err := ioutil.ReadAll(r.Body) 
		if err != nil {
			log.Fatal("Erro ao cadastrar Cliente ", err)
		}
		var novocadastro entity.Cliente
		err = json.Unmarshal(body, &novocadastro)
		if err != nil {
			log.Fatal("Erro ao cadastrar cliente", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Struct: ", novocadastro)
	} 
}

func LoginCliente(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		
	}
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


