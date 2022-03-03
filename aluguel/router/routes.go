package router

import (
	rest "github.com/LeandroAlcantara-1997/rest"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	RouterCliente(r)
	RouterAdmin(r)
	RouterVeiculo(r)
	RouterAdmin(r)
	RouterAluguel(r)

	r.HandleFunc("/recuperar", rest.PostRestauraSenha).Methods("POST")
	return r
}
