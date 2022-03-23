package router

import (
	rest "github.com/LeandroAlcantara-1997/rest"
	"github.com/LeandroAlcantara-1997/service"
	"github.com/gorilla/mux"
)

//Aluguel
func RouterAluguel(r *mux.Router){
	r.HandleFunc("/aluguel/alugados", service.MiddlewareAdmin(rest.GetAlugueis)).Methods("GET")
	r.HandleFunc("/aluguel/alugar/{cliente}/{veiculo}", service.MiddlewareCliente(rest.AlugarCarro)).Methods("POST")     
}