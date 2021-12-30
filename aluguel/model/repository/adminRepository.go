package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createAdmin(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS admin(" +
		"user VARCHAR(5) NOT NULL," +
		"senha VARCHAR(100), " +
		"token VARCHAR(200)" +
		")")

	if err != nil {
		log.Fatalf("Erro ao criar tabela Admin ", err)
	}

	err = InsertAdmin(db)
	if err != nil {
		log.Fatalf("Erro ao inserir", err)
	}

	return nil
}

func InsertAdmin(db *sql.DB) error {
	token, err := entity.GeraToken("admin" + "admin123456")
	fmt.Println(token)
	if err != nil {
		log.Fatalf("Erro ao gerar token", err)
	}
	_, err = db.Exec("INSERT INTO Admin(user, token )" +
		"VALUES ('" + "admin" +
		"' , '" + token + "')")
	if err != nil {
		log.Fatalf("Erro ao inserir admin ", err)
	}

	return nil
}

func LogarAdmin(admin *entity.Admin) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("Erro open sql", err)
	}

	fmt.Println(admin.Token)

	rows := db.QueryRow("SELECT token FROM admin " +
		"WHERE token = '" + admin.Token + "';")

	err = rows.Scan(&admin.Token)
	if err != nil {
		return fmt.Errorf("Acesso negado", err)
	}
	return nil
}
