package controller

import (
	rest "github.com/LeandroAlcantara-1997/controller/rest"
	"github.com/gorilla/mux"
)

//Cliente
func RouterCliente(r *mux.Router) {

	r.HandleFunc("/", rest.GetLoginCliente).Methods("GET") //Template
	r.HandleFunc("/", rest.PostLoginCliente).Methods("POST")

	r.HandleFunc("/cliente/cadastro", rest.GetCadastraCliente).Methods("GET") //Template
	r.HandleFunc("/cliente/cadastro", rest.PostCadastraCliente).Methods("POST")
	r.HandleFunc("/cliente/delete", rest.DeletaCadastro).Methods("DELETE")

	r.HandleFunc("/cliente/getclientebyid", rest.GetByIdCliente).Methods("GET")
	r.HandleFunc("/cliente/getclientes", rest.GetClientesCadastrados).Methods("GET")

	r.HandleFunc("/cliente/home", rest.HomeCliente).Methods("GET") //Template
}
