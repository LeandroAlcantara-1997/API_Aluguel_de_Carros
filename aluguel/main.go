package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/controller"
)

func main() {

	//r := router.PathPrefix("/api").Subrouter()

	http.HandleFunc("/aluguel", controller.GetAluguel)
	
	//Cliente
	http.HandleFunc("/cadastroCliente", controller.CadastraCliente)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem vindo")
	})

	fmt.Println("Serivdor rodando")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor", err)
	}

}
