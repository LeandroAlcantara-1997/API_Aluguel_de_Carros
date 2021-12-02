package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenSQL() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:Arthur08102019@tcp(localhost:3306)/aluguel")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	var result sql.Result
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS cliente (" +
		"id INT AUTO_INCREMENT, " +
		"nome VARCHAR(25) NOT NULL, " +
		"sobrenome VARCHAR(50) NOT NULL, " +
		"dataNascimento DATE NOT NULL, " +
		"rg	VARCHAR(10) NOT NULL UNIQUE," +
		"cpf VARCHAR(11) NOT NULL UNIQUE, " +
		"cnh VARCHAR(10) UNIQUE, " +
		"PRIMARY KEY(id)" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela clientes %v ", err)
	}
	fmt.Print(result)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS login(" +
		"email VARCHAR(70), " +
		"senha VARCHAR(100), " +
		"fk_cliente INT, " +
		"token	VARCHAR(100), " +
		"CONSTRAINT id PRIMARY KEY (fk_cliente), " +
		"FOREIGN KEY (fk_cliente) REFERENCES  cliente(id)" + ")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela login %v", err)
	}
	fmt.Println(result)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS contato(" + "id INT AUTO_INCREMENT PRIMARY KEY, " +
		"celular VARCHAR(11), " +
		"telefone VARCHAR(10), " +
		"email	VARCHAR(70), " +
		"fk_cliente INT, " +
		"FOREIGN KEY (fk_cliente) REFERENCES cliente(id)" +
		")")
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Erro ao criar tabela contato %v ", err)
	}

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS endereco(" +
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
	fmt.Println(result)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS estado(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"nome VARCHAR(25), " +
		"pais	VARCHAR(15) DEFAULT 'Brasil'" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela estado %v", err)
	}
	fmt.Println(result)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS login(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"email VARCHAR(70), " +
		"senha VARCHAR(100), " +
		"fk_cliente INT, " +
		"token	VARCHAR(100), " +
		"CONSTRAINT id PRIMARY KEY (fk_cliente), " +
		"FOREIGN KEY (fk_cliente) REFERENCES  cliente(id)" + ")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela login %v", err)
	}
	fmt.Println(result)
	return nil
}
