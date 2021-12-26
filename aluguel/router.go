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

	


	//Cliente
	http.HandleFunc("/cadastroCliente", controller.CadastraCliente)
	http.HandleFunc("/login", controller.LoginCliente)
	http.HandleFunc("/recuperarSenha", controller.RestauraSenha)
	http.HandleFunc("/homeCliente", controller.HomeCliente)

	//Carros
	http.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados)
	

	//Admin
	http.HandleFunc("/getIdCliente", controller.GetByIdCliente)
	http.HandleFunc("/getClientes", controller.GetClientesCadastrados)
	http.HandleFunc("/loginAdmin", controller.LoginAdmin)
	http.HandleFunc("/getAlugueis", controller.GetAluguel)
	http.HandleFunc("/homeAdmin", controller.HomeAdmin)

	fmt.Println("Serivdor rodando porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
