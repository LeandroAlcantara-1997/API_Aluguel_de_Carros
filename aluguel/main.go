package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/LeandroAlcantara-1997/controller/router"
	service "github.com/LeandroAlcantara-1997/controller/service"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func main() {
	r := routes.NewRouter()
	http.Handle("/", r)

	//Cria as tabelas assim que o programa é executado
	service.LoadTemplates("view/*.html")
	err := repository.CreateTables()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serivdor rodando porta 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}
