package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
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

func DeleteContato(id string) error {
	db, err := OpenSQL()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM contato " +
	"WHERE fk_cliente = '" + id + "';")
	if err != nil {
		return fmt.Errorf("Erro ao deletar contato: %#v", err)
	}

	return nil
}

func GetContatoById(id string) (entity.Contato, error) {
	var contato entity.Contato
	db, err := OpenSQL()
	if err != nil {
		return contato, fmt.Errorf("%v", err)
	}
	rows := db.QueryRow("SELECT id, celular, telefone, email FROM contato " + "WHERE fk_cliente= " + id + ";")
	err = rows.Scan(&contato.Id, &contato.Celular, &contato.Telefone, &contato.Email)
	if err != nil {
		return contato, fmt.Errorf("Erro ao pegar dados do contato %#v", err)
	}
	return contato, nil
}