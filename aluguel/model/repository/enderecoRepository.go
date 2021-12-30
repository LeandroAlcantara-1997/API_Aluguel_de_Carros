package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createEndereco(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS endereco(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"fk_estado INT, " +
		"cidade VARCHAR(25), " +
		"bairro VARCHAR(25), " +
		"logradouro	VARCHAR(15), " +
		"rua	VARCHAR(50), " +
		"numero	VARCHAR(10), " +
		"complemento VARCHAR(15), " +
		"fk_cliente INT, " +
		"FOREIGN KEY (fk_cliente) REFERENCES cliente(id)" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela endereco %v", err)
	}

	return nil
}

func InsertEndereco(c *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}
	idEstado, err := entity.ValidaEstado(c.Endereco.Estado.Nome)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	_, err = db.Exec("INSERT INTO endereco (fk_estado, cidade, bairro, logradouro, rua, numero, complemento, fk_cliente) VALUES ('" + fmt.Sprint(idEstado) + "', '" + c.Endereco.Cidade + "', '" + c.Endereco.Bairro + "', '" + c.Endereco.Logradouro + "', '" + c.Endereco.Rua + "', '" + c.Endereco.Numero + "', '" + c.Endereco.Complemento + "', '" + fmt.Sprint(c.Id) + "')")

	if err != nil {
		return fmt.Errorf("Error ao fazer insert endereco %v", err)
	}

	err = InsertEstado(c.Endereco)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}
	return nil
}
