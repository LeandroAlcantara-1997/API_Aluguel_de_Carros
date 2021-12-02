package repository

import (
	"database/sql"
	"fmt"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func InsertCliente(cliente *entity.Cliente) (sql.Result, error) {
	db, err := OpenSQL()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	fmt.Println(cliente.Data_Nascimento)
	result, err := db.Exec("INSERT INTO cliente (nome, sobrenome, dataNascimento, rg, cpf, cnh) VALUES ('" + cliente.Nome + "', '" + cliente.Sobrenome + "', '" + cliente.Data_Nascimento + "', '" +  cliente.RG + "', '" + cliente.CPF + "', '" + cliente.CNH + "')")
	if err != nil {
		return nil, fmt.Errorf("Error ao inserir dados na tabela cliente %v", err)
	}

	return result, nil
}

func InsertLogin(c *entity.Cliente) (sql.Result, error) {
	db, err := OpenSQL()

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	result, err := db.Exec("INSERT INTO login (email, senha, fk_cliente, token) VALUES ('" + c.Contato.Email + "', '" + c.Login.Senha + "', '" + string(c.Id) + "', '" + c.Login.Token + "')")

	if err != nil {
		return nil, fmt.Errorf("Error ao fazer insert do login %v", err)
	}

	return result, nil
}