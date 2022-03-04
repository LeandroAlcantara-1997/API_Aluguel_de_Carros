package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenSQL() (*sql.DB, error) {
	var db *sql.DB
														   //mysql-db
	db, err := sql.Open("mysql", "root@tcp(mysql-db:3306)/aluguel_veiculo")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func CreateTables() error {
	db, err := OpenSQL()
	if err  != nil {
		log.Fatalf("%v", err)
	}
	err = createCliente(db)
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
	if err != nil {
		log.Fatal(err)
	}

	err = createVeiculos(db)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}






