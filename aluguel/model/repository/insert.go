package repository

import (
	"database/sql"
	"fmt"
	"log"

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

	result, err = InsertContato(cliente)
	rowsAffected, err = result.RowsAffected()
	if  err != nil {
		fmt.Println("Erro ao retornar número de linhas afetadas", err)
	}
	fmt.Println("Linhas afetadas no insert contato", rowsAffected)
	fmt.Println("Erro começa aqui")

	result, err = InsertEndereco(cliente)
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao retornar número de linhas afetadas", err)
	}
	fmt.Println("Linhas afetadas no insert endereco", rowsAffected)
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

func InsertContato(c *entity.Cliente) (sql.Result, error){
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO contato (celular, telefone, email, fk_cliente) VALUES ('" + c.Contato.Celular + "', '" + c.Contato.Telefone + "', '" + c.Contato.Email + "', '" + fmt.Sprint(c.Id) + "')")
	if err != nil {
		return nil, fmt.Errorf("Error ao fazer insert contato %v", err)
	}
	return result, nil
}

func InsertEndereco(c *entity.Cliente) (sql.Result, error) {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}
	idEstado, err := entity.ValidaEstado(c.Endereco.Estado)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	result, err := db.Exec("INSERT INTO endereco (fk_estado, cidade, bairro, logradouro, rua, numero, complemento, fk_cliente) VALUES ('" + fmt.Sprint(idEstado) + "', '" + c.Endereco.Cidade + "', '" + c.Endereco.Bairro + "', '" + c.Endereco.Logradouro + "', '" + c.Endereco.Rua + "', '" + c.Endereco.Numero + "', '" + c.Endereco.Complemento + "', '" + fmt.Sprint(c.Id) + "')" )

	if err != nil {
		return nil, fmt.Errorf("Error ao fazer insert endereco %v", err)
	}
	return result, nil
}