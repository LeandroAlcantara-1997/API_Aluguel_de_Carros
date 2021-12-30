package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createEstado(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS estado(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"nome VARCHAR(25) UNIQUE, " +
		"pais	VARCHAR(15) DEFAULT 'Brasil'" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela estado %v", err)
	}

	return nil
}

func InsertEstado(endereco entity.Endereco) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO estado (nome, pais) " +
		"VALUES ('" + endereco.Estado.Nome + "', '" + endereco.Estado.Pais + "');")
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela estado")
	}

	return nil
}
