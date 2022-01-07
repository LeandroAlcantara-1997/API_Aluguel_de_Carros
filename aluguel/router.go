package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/controller"
	utils "github.com/LeandroAlcantara-1997/controller/utils"
	"github.com/LeandroAlcantara-1997/model/repository"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./view/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	http.Handle("/", r)

	//Cria as tabelas assim que o programa é executado
	utils.LoadTemplates("view/*.html")
	err := repository.CreateTables()
	if err != nil {
		log.Fatal(err)
	}

	//Cliente
	r.HandleFunc("/cadastroCliente", controller.GetCadastraCliente).Methods("GET") //Template
	r.HandleFunc("/cadastroCliente", controller.PostCadastraCliente).Methods("POST")

	r.HandleFunc("/deletaCadastro", controller.DeletaCadastro).Methods("DELETE")

	r.HandleFunc("/login", controller.GetLoginCliente).Methods("GET") //Template
	r.HandleFunc("/login", controller.PostLoginCliente).Methods("POST")

	r.HandleFunc("/recuperarSenha", controller.GetRestauraSenha).Methods("GET") //Template
	r.HandleFunc("/recuperarSenha", controller.PostRestauraSenha).Methods("POST")

	r.HandleFunc("/homeCliente", controller.HomeCliente) //Template

	//Carros
	r.HandleFunc("/carrosCadastrados", controller.GetCarrosCadastrados)
	r.HandleFunc("/cadastraCarro", controller.CadastraCarro)
	//r.HandleFunc("/carrosDisponiveis", controller.CarrosDisponiveis)

	//Admin
	r.HandleFunc("/getIdCliente", controller.GetByIdCliente)
	r.HandleFunc("/getClientes", controller.GetClientesCadastrados)

	r.HandleFunc("/loginAdmin", controller.GetLoginAdmin).Methods("GET") //Template
	r.HandleFunc("/loginAdmin", controller.PostLoginAdmin).Methods("POST")

	r.HandleFunc("/homeAdmin", controller.HomeAdmin)

	//Aluguel
	r.HandleFunc("/carrosAlugados", controller.GetAlugueis).Methods("GET") //Template
	r.HandleFunc("/alugar", controller.AlugarCarro)                        //Template

	fmt.Println("Serivdor rodando porta 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
