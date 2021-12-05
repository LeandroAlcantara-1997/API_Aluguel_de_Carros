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
	route.HandleFunc("/", controller.LoginCliente).Methods("GET")
	route.HandleFunc("/recuperarSenha", controller.RestauraSenha).Methods("GET")

	//Carros
	route.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados).Methods("GET")

	//Admin
	route.HandleFunc("/getIdCliente/{id}", controller.GetByIdCliente).Methods("GET")
	route.HandleFunc("/getClientes", controller.GetClientesCadastrados).Methods("GET")
	route.HandleFunc("/loginAdmin", controller.LoginAdmin)

	fmt.Println("Serivdor rodando porta 8080")
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
