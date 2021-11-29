package repository

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func OpenSQL(){
	db, err := sql.Open("mysql", "root:Arthur08102019@tcp(localhost:3306)/aluguel")
	if err != nil {
		fmt.Errorf("Erro ao configurar Database %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}