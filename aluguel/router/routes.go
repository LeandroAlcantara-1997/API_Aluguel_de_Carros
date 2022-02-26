package router

import (
	"net/http"

	rest "github.com/LeandroAlcantara-1997/rest"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./view/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	RouterCliente(r)
	RouterAdmin(r)
	RouterVeiculo(r)
	RouterAdmin(r)
	RouterAluguel(r)

	r.HandleFunc("/recuperar", rest.PostRestauraSenha).Methods("POST")
	return r
}
