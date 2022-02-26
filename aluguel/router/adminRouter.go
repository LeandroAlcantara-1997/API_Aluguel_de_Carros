package router

import (
	"github.com/gorilla/mux"
	rest "github.com/LeandroAlcantara-1997/rest"
)
//Admin
func RouterAdmin(r *mux.Router) {
	r.HandleFunc("/admin/login", rest.PostLoginAdmin).Methods("POST")
}