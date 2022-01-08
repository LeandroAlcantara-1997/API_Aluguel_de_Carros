package controller

import (
	"github.com/gorilla/mux"
	rest "github.com/LeandroAlcantara-1997/controller/rest"
)
//Admin
func RouterAdmin(r *mux.Router) {
	r.HandleFunc("/admin/login", rest.GetLoginAdmin).Methods("GET") //Template
	r.HandleFunc("/admin/login", rest.PostLoginAdmin).Methods("POST")
	r.HandleFunc("/admin/home", rest.HomeAdmin).Methods("GET")
}