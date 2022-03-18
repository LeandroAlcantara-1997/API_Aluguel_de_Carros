package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/repository"
	routes "github.com/LeandroAlcantara-1997/router"
)

func main() {
	r := routes.NewRouter()

	//Apaga as tabelas se existirem
	err := repository.DropTables()
	if err != nil {
		log.Fatal(err)
	}
	
	//Cria as tabelas assim que o programa é executado
	err = repository.CreateTables()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serivdor rodando porta 8081")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Erro ao rodar servidor ", err)
	}

}

func soma(i ...int) int {
	var total int
	for _, v := range i {
		total += v
	}
	return total
}
