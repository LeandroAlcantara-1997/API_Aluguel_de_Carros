package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func GetByIdCliente(id int) (entity.Cliente, error) {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal("%v", err)
	}
	var rows *sql.Rows
	var cliente entity.Cliente
	rows, err = db.Query("SELECT nome, sobrenome, dataNascimento, rg, cpf, cnh FROM 					cliente " + "WHERE id=" + fmt.Sprint(id))
	if err != nil {
		log.Fatal("Erro ao retornar cadastro", err)
	}

	for rows.Next(){
		err = rows.Scan(&cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
		if err != nil {
			log.Fatal("Erro ao pegar dados do cliente ", err)
		}
	}
	return cliente, nil
}