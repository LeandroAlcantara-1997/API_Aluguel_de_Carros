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
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./view/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	http.Handle("/", r)

	//Cria as tabelas assim que o programa é executado
	controller.LoadTemplates("view/*.html")
	err := repository.CreateTables()
	if err != nil {
		log.Fatal(err)
	}


	//Cliente
	r.HandleFunc("/cadastroCliente", controller.GetCadastraCliente).Methods("GET")
	r.HandleFunc("/cadastroCliente", controller.PostCadastraCliente).Methods("POST")
	
	r.HandleFunc("/login", controller.GetLoginCliente).Methods("GET")
	r.HandleFunc("/login", controller.PostLoginCliente).Methods("POST")
	
	r.HandleFunc("/recuperarSenha", controller.GetRestauraSenha).Methods("GET")
	r.HandleFunc("/recuperarSenha", controller.PostRestauraSenha).Methods("POST")
	
	r.HandleFunc("/homeCliente", controller.HomeCliente)

	//Carros
	r.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados)

	//Admin
	r.HandleFunc("/getIdCliente", controller.GetByIdCliente)
	r.HandleFunc("/getClientes", controller.GetClientesCadastrados)

	r.HandleFunc("/loginAdmin", controller.GetLoginAdmin).Methods("GET")
	r.HandleFunc("/loginAdmin", controller.PostLoginAdmin).Methods("POST")
	
	r.HandleFunc("/getAlugueis", controller.GetAluguel)
	r.HandleFunc("/homeAdmin", controller.HomeAdmin)

	r.HandleFunc("/cadastraCarro", controller.CadastraCarro)

	fmt.Println("Serivdor rodando porta 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
