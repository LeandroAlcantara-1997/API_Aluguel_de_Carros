package repository

import (
	"database/sql"
	"fmt"
)

func GetByIdCliente(id int) (*sql.Rows, error) {
	var db *sql.DB
	var rows *sql.Rows
	rows, err := db.Query("SELECT * FROM cliente " +
		"WHERE id==" + fmt.Sprint(id) +
		")")
	if err != nil {
		return nil, fmt.Errorf("Erro ao retornar cadastro")
	}
	return rows, nil
}