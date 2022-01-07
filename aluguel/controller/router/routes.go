package controller

import (
	"net/http"

	rest "github.com/LeandroAlcantara-1997/controller/rest"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./view/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	//Cliente
	r.HandleFunc("/cliente/cadastro", rest.GetCadastraCliente).Methods("GET") //Template
	r.HandleFunc("/cliente/cadastro", rest.PostCadastraCliente).Methods("POST")
	r.HandleFunc("/cliente/delete", rest.DeletaCadastro).Methods("DELETE")
	r.HandleFunc("/", rest.GetLoginCliente).Methods("GET") //Template
	r.HandleFunc("/", rest.PostLoginCliente).Methods("POST")

	r.HandleFunc("/recuperarSenha", rest.GetRestauraSenha).Methods("GET") //Template
	r.HandleFunc("/recuperarSenha", rest.PostRestauraSenha).Methods("POST")

	r.HandleFunc("/homeCliente", rest.HomeCliente) //Template

	//Carros
	r.HandleFunc("/carrosCadastrados", rest.GetCarrosCadastrados)
	r.HandleFunc("/cadastraCarro", rest.CadastraCarro)
	//r.HandleFunc("/carrosDisponiveis", rest.CarrosDisponiveis)

	//Admin
	r.HandleFunc("admin/loginAdmin", rest.GetLoginAdmin).Methods("GET") //Template
	r.HandleFunc("admin/loginAdmin", rest.PostLoginAdmin).Methods("POST")
	r.HandleFunc("admin/home", rest.HomeAdmin)
	r.HandleFunc("/getIdCliente", rest.GetByIdCliente)
	r.HandleFunc("/getClientes", rest.GetClientesCadastrados)

	

	//Aluguel
	r.HandleFunc("/carrosAlugados", rest.GetAlugueis).Methods("GET") //Template
	r.HandleFunc("/alugar", rest.AlugarCarro)                        //Template
	return r
}
