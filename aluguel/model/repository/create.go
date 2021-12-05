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
	err = CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	err := createCliente(db)
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	err = createLogin(db)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = createContato(db)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = createEndereco(db)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = createEstado(db)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = createAdmin(db)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = createAluguel(db) 
	if err != nil{
		log.Fatal(err)
	}

	err = createVeiculos(db) 
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func createCliente(db *sql.DB) error {
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
		log.Fatalf("Erro ao criar tabela cliente %v ", err)
	}
	fmt.Print(result.RowsAffected())
	return nil
}

func createLogin(db *sql.DB) error {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS login(" +
		"email VARCHAR(70), " +
		"senha VARCHAR(100), " +
		"fk_cliente INT, " +
		"token	VARCHAR(100), " +
		"CONSTRAINT id PRIMARY KEY (fk_cliente), " +
		"FOREIGN KEY (fk_cliente) REFERENCES  cliente(id)" + ")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela login %v", err)
	}

	fmt.Println(result.RowsAffected())

	return nil
}

func createContato(db *sql.DB) error{
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS contato(" + "id INT AUTO_INCREMENT PRIMARY KEY, " +
		"celular VARCHAR(11), " +
		"telefone VARCHAR(10), " +
		"email	VARCHAR(70), " +
		"fk_cliente INT, " +
		"FOREIGN KEY (fk_cliente) REFERENCES cliente(id)" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela contato %v ", err)
	}
	fmt.Println(result.RowsAffected())
	return nil
}

func createEndereco(db *sql.DB) error {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS endereco(" +
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

	fmt.Println(result.RowsAffected())
	return nil

}

func createEstado(db *sql.DB) error{
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS estado(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"nome VARCHAR(25) UNIQUE, " +
		"pais	VARCHAR(15) DEFAULT 'Brasil'" +
		")")
	if err != nil {
		log.Fatalf("Erro ao criar tabela estado %v", err)
	}
	fmt.Println(result.RowsAffected())
	return nil
}

func createAdmin(db *sql.DB) error {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS admin(" + 
		"user VARCHAR(5) NOT NULL," + 
		"senha VARCHAR(100), " +
		"token VARCHAR(200)" + 
		")")

		if err != nil {
			log.Fatalf("Erro ao criar tabela Admin ", err)
		}

		fmt.Println(result.RowsAffected())
		return nil
}

func createAluguel(db *sql.DB) error {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS aluguel(" +
		"data_inicio TIMESTAMP, " + 
		"data_retorno TIMESTAMP, " + 
		"total DOUBLE, " + 
		"fk_veiculo INT, " + 
		"fk_cliente INT" +
		")")

		if err != nil {
			log.Fatalf("Erro ao criar tabela aluguel ", err)
		}
		fmt.Println(result.RowsAffected())

		return nil
}

func createVeiculos(db *sql.DB) error {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS veiculo(" + 
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"modelo VARCHAR(50), " +
		"marca	VARCHAR(50), " +
		"ano YEAR," +
		"cor VARCHAR(20), " +
		"km_litro FLOAT, " + 
		"valor_dia DOUBLE," +
		"valor_hora DOUBLE" +
		")" )
	if err != nil {
		log.Fatalf("Erro ao criar tabela veiculo")
	}
	fmt.Println(result.RowsAffected())
	return nil
}