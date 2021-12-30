package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createContato(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS contato(" + "id INT AUTO_INCREMENT PRIMARY KEY, " +
		"celular VARCHAR(11), " +
		"telefone VARCHAR(10), " +
		"email	VARCHAR(70), " +
		"fk_cliente INT, " +
		"FOREIGN KEY (fk_cliente) REFERENCES cliente(id)" +
		");")
	if err != nil {
		log.Fatalf("Erro ao criar tabela contato %v ", err)
	}
	return nil
}

func InsertContato(c *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO contato (celular, telefone, email, fk_cliente) VALUES ('" + c.Contato.Celular + "', '" + c.Contato.Telefone + "', '" + c.Contato.Email + "', '" + fmt.Sprint(c.Id) + "');")
	if err != nil {
		return fmt.Errorf("Error ao fazer insert contato %v", err)
	}
	return nil
}
