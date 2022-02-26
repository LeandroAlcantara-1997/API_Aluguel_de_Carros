package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
)

func createLogin(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS login(" +
		"email VARCHAR(70), " +
		"fk_cliente INT, " +
		"token	VARCHAR(100), " +
		"CONSTRAINT id PRIMARY KEY (fk_cliente), " +
		"FOREIGN KEY (fk_cliente) REFERENCES  cliente(id)" + ");")
	if err != nil {
		log.Fatalf("Erro ao criar tabela login %v", err)
	}

	return nil
}

func InsertLogin(c *entity.Cliente) error {
	db, err := OpenSQL()

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = db.Exec("INSERT INTO login (email, fk_cliente, token) VALUES ('" + c.Contato.Email + "', '" + fmt.Sprint(c.Id) + "', '" + c.Login.Token + "');")

	if err != nil {
		return fmt.Errorf("Error ao fazer insert do login %v", err)
	}
	return nil
}

func GetEmailToSenha(email string) error {
	db, err := OpenSQL()
	var log entity.Login
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	rows := db.QueryRow("SELECT email FROM login " +
		"WHERE email='" + email + "';")
	fmt.Print(rows)
	err = rows.Scan(&log.Email)
	if err != nil {
		return fmt.Errorf("Email não cadastrado", err)
	}

	return nil
}

func Logar(email, senha string) error {
	var log entity.Login

	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	token, err := entity.GeraToken(email + senha)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}
	rows := db.QueryRow("SELECT token FROM login " +
		"WHERE token='" + token + "';")

	err = rows.Scan(&log.Token)
	if err != nil {
		return fmt.Errorf("Acesso negado", err)
	}

	return nil
}

func UpdateSenha(email, senha string) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("Erro ao atualizar senha: %v", err)
	}
	token, err := entity.GeraToken(email + senha)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	_, err = db.Exec("UPDATE login " +
		"SET token='" + token +
		"' WHERE email = '" + email + "';")
	if err != nil {
		return fmt.Errorf("Erro ao fazer update %v", err)
	}
	return nil
}

func DeleteLogin(id string) error {
	db, err := OpenSQL()
	if err != nil {
		return err
	}

	_ , err = db.Exec("DELETE FROM login " + 
	"WHERE fk_cliente = '" + id + "';")

	if err != nil {
		return fmt.Errorf("Erro ao apagar login", err)
	}

	return nil
}