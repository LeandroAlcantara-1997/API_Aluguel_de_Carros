package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
)

func createAdmin(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS admin(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"user VARCHAR(5) NOT NULL," +
		"senha VARCHAR(200)" +
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
	_, err = db.Exec("INSERT INTO admin(user, senha )" +
		"VALUES ('" + "admin" +
		"' , '" + token + "')")
	if err != nil {
		log.Fatalf("Erro ao inserir admin %#v", err)
	}

	return nil
}

func LogarAdmin(admin *entity.Admin) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("Erro open sql %#v", err)
	}

	rows := db.QueryRow("SELECT senha FROM admin " +
		"WHERE senha = '" + admin.Senha + "';")

	if err = rows.Scan(&admin.Senha); err != nil {
		return fmt.Errorf("Acesso negado %#v", err)
	}

	return nil
}

func GetIdAdminByUser(user string) (int64, error) {
	var id int64
	db, err := OpenSQL()
	if err != nil {
		return 0, fmt.Errorf("Erro open sql %#v", err)
	}

	rows := db.QueryRow("SELECT id FROM admin " +
		"WHERE user = '" + user + "';")

	if err = rows.Scan(&id); err != nil {
		return 0, fmt.Errorf("Acesso negado %#v", err)
	}

	return id, nil
}
