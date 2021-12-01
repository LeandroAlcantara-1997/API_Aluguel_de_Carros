package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/controller"
	"github.com/LeandroAlcantara-1997/model/repository"
	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/aluguel", controller.GetAluguel).Methods("GET")

	//Cliente
	route.HandleFunc("/cadastroCliente", controller.CadastraCliente).Methods("POST")
	route.HandleFunc("/loginCliente", controller.LoginCliente).Methods("GET")

	//Carros
	route.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados).Methods("GET")

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem vindo")
		repository.OpenSQL()
		repository.CreateTable()
		return
	})

	fmt.Println("Serivdor rodando porta 8080")
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
