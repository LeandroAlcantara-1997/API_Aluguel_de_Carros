package router

import (
	"github.com/gorilla/mux"
	rest "github.com/LeandroAlcantara-1997/rest"
)

//Veiculos
func RouterVeiculo(r *mux.Router) {
	
	r.HandleFunc("/veiculo/cadastrados", rest.GetCarrosCadastrados).Methods("GET")
	r.HandleFunc("/veiculo/cadastro", rest.CadastraCarro).Methods("POST")
	//r.HandleFunc("veiculo/disponiveis", rest.CarrosDisponiveis)
}