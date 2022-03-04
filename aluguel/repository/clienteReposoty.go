package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
)

func createCliente(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cliente (" +
		"id INT AUTO_INCREMENT, " +
		"nome VARCHAR(25) NOT NULL, " +
		"sobrenome VARCHAR(50) NOT NULL, " +
		"dataNascimento DATE NOT NULL, " +
		"rg	VARCHAR(10) NOT NULL UNIQUE," +
		"cpf VARCHAR(11) NOT NULL UNIQUE, " +
		"cnh VARCHAR(10) UNIQUE, " +
		"PRIMARY KEY(id)" +
		");")
	if err != nil {
		log.Fatalf("Erro ao criar tabela cliente %v ", err)
	}
	return nil
}

func InsertCliente(cliente *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	result, err := db.Exec("INSERT INTO cliente (nome, sobrenome, dataNascimento, rg, cpf, cnh) VALUES ('" + cliente.Nome + "', '" + cliente.Sobrenome + "', '" + cliente.Data_Nascimento + "', '" + cliente.RG + "', '" + cliente.CPF + "', '" + cliente.CNH + "');")
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela cliente %v", err)
	}

	cliente.Id, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Erro ao recuperar ultimo id %#v", err)
	}

	err = InsertLogin(cliente)
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela login %v", err)
	}

	err = InsertContato(cliente)
	if err != nil {
		return fmt.Errorf("Erro na funcao InsertContato: %#v", err)
	}

	err = InsertEndereco(cliente)
	if err != nil {
		return fmt.Errorf("Erro na funcao InsertContato: %#v", err)
	}

	return nil
}

func GetClientesCadastrados() ([]entity.Cliente, error) {
	var cliente entity.Cliente
	var clientes []entity.Cliente
	db, err := OpenSQL()
	if err != nil {
		return nil, fmt.Errorf("%#v", err)
	}

	rows, err := db.Query("SELECT id, nome, sobrenome, dataNascimento, rg, cpf, cnh FROM cliente;")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select na tabela cleinte: %#v", err)
	}

	for rows.Next() {
		err = rows.Scan(&cliente.Id, &cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
		if err != nil {
			return nil, fmt.Errorf("Erro ao passar valores para cliente")
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func GetByIdCliente(id int) (entity.Cliente, error) {
	var cliente entity.Cliente

	db, err := OpenSQL()
	if err != nil {
		return cliente, fmt.Errorf("%v", err)
	}

	rows := db.QueryRow("SELECT nome, sobrenome, dataNascimento, rg, cpf, cnh FROM 	cliente " + "WHERE id= " + fmt.Sprint(id) + ";")

	err = rows.Scan(&cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
	if err != nil {
		return cliente, fmt.Errorf("Erro ao pegar dados do cliente %#v", err)
	}

	return cliente, nil
}

func DeletaCliente(id string) (error) {
	if err := DeleteAluguel(id); err != nil {
		return err
	}

	if err := DeleteEndereco(id); err != nil {
		return err
	}

	if err := DeleteContato(id); err != nil {
		return err
	}

	if err := DeleteLogin(id); err != nil {
		return err
	}

	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_ , err = db.Exec("DELETE FROM cliente " +
	"WHERE id = '" + id + "';")
	if err != nil {
		return fmt.Errorf("Erro ao deletar cadastro %#v", err)
	}

	return nil
}