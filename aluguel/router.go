package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/controller"
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

	//Admin
	route.HandleFunc("/getIdCliente/{id}" , controller.GetByIdCliente).Methods("GET")

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem vindo")
	})

	fmt.Println("Serivdor rodando porta 8080")
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
