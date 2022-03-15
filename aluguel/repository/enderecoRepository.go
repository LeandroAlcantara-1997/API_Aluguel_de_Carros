package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
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
		");")
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
	c.Endereco.Estado.Id, err = entity.ValidaEstado(c.Endereco.Estado.Nome)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	_, err = db.Exec("INSERT INTO endereco (fk_estado, cidade, bairro, logradouro, rua, numero, complemento, fk_cliente) VALUES ('" + fmt.Sprint(c.Endereco.Estado.Id) + "', '" + c.Endereco.Cidade + "', '" + c.Endereco.Bairro + "', '" + c.Endereco.Logradouro + "', '" + c.Endereco.Rua + "', '" + c.Endereco.Numero + "', '" + c.Endereco.Complemento + "', '" + fmt.Sprint(c.Id) + "');")
	if err != nil {
		return fmt.Errorf("Error ao fazer insert endereco %v", err)
	}

	err = InsertEstado(&c.Endereco)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}
	return nil
}

func DeleteEndereco(id string) error {
	db, err := OpenSQL()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM endereco " +
		"WHERE fk_cliente = '" + id + "';")
	if err != nil {
		return fmt.Errorf("Erro ao deletar endereco %#v", err)
	}

	return nil
}

func GetEnderecoById(id string) (entity.Endereco, error) {
	var endereco entity.Endereco

	db, err := OpenSQL()
	if err != nil {
		return endereco, fmt.Errorf("%v", err)
	}
	rows := db.QueryRow("SELECT id, fk_estado, cidade, bairro, logradouro, rua, numero, complemento FROM endereco " + "WHERE fk_cliente= " + id + ";")

	err = rows.Scan(&endereco.Id, &endereco.Estado.Id, &endereco.Cidade, &endereco.Bairro, &endereco.Logradouro, &endereco.Rua, &endereco.Numero, &endereco.Complemento)
	if err != nil {
		return endereco, fmt.Errorf("Erro ao pegar dados do endereco %#v", err)
	}

	return endereco, nil
}
