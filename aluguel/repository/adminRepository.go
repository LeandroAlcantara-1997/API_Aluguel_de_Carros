package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
)

func createAdmin(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS admin(" +
		"user VARCHAR(5) NOT NULL," +
		"senha VARCHAR(100), " +
		"token VARCHAR(200)" +
		")"); err != nil {
			log.Fatalf("Erro ao criar tabela Admin %#v", err)
		}

	if err := InsertAdmin(db); err != nil {
		log.Fatalf("Erro ao inserir %#v", err)
	}

	return nil
}

func InsertAdmin(db *sql.DB) error {
	token, err := entity.GeraToken("admin" + "admin123456")
	if err != nil {
		log.Fatalf("Erro ao gerar token %#v", err)
	}
	_, err = db.Exec("INSERT INTO admin(user, token )" +
		"VALUES ('" + "admin" +
		"' , '" + token + "')")
	if err != nil {
		log.Fatalf("Erro ao inserir admin %#v", err)
	}

	return nil
}

func LogarAdmin(admin *entity.Admin) error {
	db, err := OpenSQL(); 
	if err != nil {
		return fmt.Errorf("Erro open sql %#v", err)
	}

	rows := db.QueryRow("SELECT token FROM admin " +
		"WHERE token = '" + admin.Token + "';")

	if err = rows.Scan(&admin.Token); err != nil {
		return fmt.Errorf("Acesso negado %#v", err)
	}

	return nil
}
