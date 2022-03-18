package router

import (
	rest "github.com/LeandroAlcantara-1997/rest"
	"github.com/LeandroAlcantara-1997/service"
	"github.com/gorilla/mux"
)

//Cliente
func RouterCliente(r *mux.Router) {

	r.HandleFunc("/", rest.PostLoginCliente).Methods("POST")
	
	r.HandleFunc("/cliente/cadastro", rest.PostCadastraCliente).Methods("POST")
	r.HandleFunc("/cliente/delete/{id}", service.MiddlewareCliente(rest.DeletaCadastro)).Methods("DELETE")

	r.HandleFunc("/cliente/getclientebyid/{id}", service.MiddlewareCliente(rest.GetClienteById)).Methods("GET")
	r.HandleFunc("/cliente/getclientes", rest.GetClientesCadastrados).Methods("GET")
	r.HandleFunc("/cliente/alugados/{id}", service.MiddlewareCliente(rest.GetCarrosAlugados)).Methods("GET")
}
