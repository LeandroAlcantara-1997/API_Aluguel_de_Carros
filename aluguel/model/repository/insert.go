package repository

import (
	"database/sql"
	"fmt"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func InsertCliente(cliente *entity.Cliente) (error) {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	result, err := db.Exec("INSERT INTO cliente (nome, sobrenome, dataNascimento, rg, cpf, cnh) VALUES ('" + cliente.Nome + "', '" + cliente.Sobrenome + "', '" + cliente.Data_Nascimento + "', '" +  cliente.RG + "', '" + cliente.CPF + "', '" + cliente.CNH + "')")
	if err != nil {
		return fmt.Errorf("Error ao inserir dados na tabela cliente %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if  err != nil {
		fmt.Println("Erro ao retornar número de linhas afetadas", err)
	}
	fmt.Println("Linhas afetadas no insert cliente", rowsAffected)
	
	result, err = InsertLogin(cliente)
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela login %v", err)
	}

	rowsAffected, err = result.RowsAffected()
	if  err != nil {
		fmt.Println("Erro ao retornar número de linhas afetadas", err)
	}
	fmt.Println("Linhas afetadas no insert login", rowsAffected)

	return nil
}

func InsertLogin(c *entity.Cliente) (sql.Result, error) {
	db, err := OpenSQL()

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	result, err := db.Exec("INSERT INTO login (email, senha, fk_cliente, token) VALUES ('" + c.Contato.Email + "', '" + c.Login.Senha + "', '" + fmt.Sprint(c.Id) + "', '" + c.Login.Token + "')")

	if err != nil {
		return nil, fmt.Errorf("Error ao fazer insert do login %v", err)
	}

	return result, nil
}