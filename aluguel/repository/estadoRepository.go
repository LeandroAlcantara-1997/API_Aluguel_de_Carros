package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/LeandroAlcantara-1997/entity"
)

func createEstado(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS estado(" +
		"id INT PRIMARY KEY, " +
		"nome VARCHAR(25), " +
		"pais	VARCHAR(15) DEFAULT 'Brasil'" +
		");")
	if err != nil {
		log.Fatalf("Erro ao criar tabela estado %v", err)
	}

	return nil
}

func InsertEstado(endereco *entity.Endereco) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}
	
	_ , err = db.Exec("INSERT IGNORE INTO estado (id, nome, pais) " +
		"VALUES ('" + strconv.Itoa(endereco.Estado.Id) + "', '" + endereco.Estado.Nome + "', '" + endereco.Estado.Pais + "');")
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela estado %#v", err)
	}
	
	return nil
}

func GetEstadoById(id string) (entity.Estado, error) {
	var estado entity.Estado

	db, err := OpenSQL()
	if err != nil {
		return estado, fmt.Errorf("%v", err)
	}
	rows := db.QueryRow("SELECT id, nome, pais FROM estado " + "WHERE id= " + id + ";")
	err = rows.Scan(&estado.Id, &estado.Nome, &estado.Pais)
	if err != nil {
		return estado, fmt.Errorf("Erro ao pegar dados do estado %#v", err)
	}
	return estado, nil
}
