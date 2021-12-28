package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func InsertCliente(cliente *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = db.Exec("INSERT INTO cliente (nome, sobrenome, dataNascimento, rg, cpf, cnh) VALUES ('" + cliente.Nome + "', '" + cliente.Sobrenome + "', '" + cliente.Data_Nascimento + "', '" + cliente.RG + "', '" + cliente.CPF + "', '" + cliente.CNH + "')")
	if err != nil {
		return fmt.Errorf("Erro ao inserir dados na tabela cliente %v", err)
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

func InsertLogin(c *entity.Cliente) error {
	db, err := OpenSQL()

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = db.Exec("INSERT INTO login (email, fk_cliente, token) VALUES ('" + c.Contato.Email + "', '" + fmt.Sprint(c.Id) + "', '" + c.Login.Token + "')")

	if err != nil {
		return fmt.Errorf("Error ao fazer insert do login %v", err)
	}
	return nil
}

func InsertContato(c *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO contato (celular, telefone, email, fk_cliente) VALUES ('" + c.Contato.Celular + "', '" + c.Contato.Telefone + "', '" + c.Contato.Email + "', '" + fmt.Sprint(c.Id) + "')")
	if err != nil {
		return fmt.Errorf("Error ao fazer insert contato %v", err)
	}
	return nil
}

func InsertEndereco(c *entity.Cliente) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}
	idEstado, err := entity.ValidaEstado(c.Endereco.Estado)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	_, err = db.Exec("INSERT INTO endereco (fk_estado, cidade, bairro, logradouro, rua, numero, complemento, fk_cliente) VALUES ('" + fmt.Sprint(idEstado) + "', '" + c.Endereco.Cidade + "', '" + c.Endereco.Bairro + "', '" + c.Endereco.Logradouro + "', '" + c.Endereco.Rua + "', '" + c.Endereco.Numero + "', '" + c.Endereco.Complemento + "', '" + fmt.Sprint(c.Id) + "')")

	if err != nil {
		return fmt.Errorf("Error ao fazer insert endereco %v", err)
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
