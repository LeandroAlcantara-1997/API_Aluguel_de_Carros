package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/controller"
)

func main() {

	fs := http.FileServer(http.Dir("./view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))


	http.HandleFunc("/aluguel", controller.GetAluguel)

	//Cliente
	http.HandleFunc("/cadastroCliente", controller.CadastraCliente)
	http.HandleFunc("/", controller.LoginCliente)
	http.HandleFunc("/recuperarSenha", controller.RestauraSenha)

	//Carros
	http.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados)

	//Admin
	http.HandleFunc("/getIdCliente/{id}", controller.GetByIdCliente)
	http.HandleFunc("/getClientes", controller.GetClientesCadastrados)
	http.HandleFunc("/loginAdmin", controller.LoginAdmin)

	fmt.Println("Serivdor rodando porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
