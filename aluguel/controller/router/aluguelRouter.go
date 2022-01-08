package controller


import (
	"github.com/gorilla/mux"
	rest "github.com/LeandroAlcantara-1997/controller/rest"
)

//Aluguel
func RouterAluguel(r *mux.Router){
	r.HandleFunc("/aluguel/alugados", rest.GetAlugueis).Methods("GET") //Template
	r.HandleFunc("/aluguel/alugar", rest.AlugarCarro).Methods("POST")     
}