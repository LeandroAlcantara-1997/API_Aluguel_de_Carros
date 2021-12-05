package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func GetByIdCliente(id int) (entity.Cliente, error) {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal("%v", err)
	}
	var rows *sql.Rows
	var cliente entity.Cliente
	//Alterar Query para QueryRow
	rows, err = db.Query("SELECT nome, sobrenome, dataNascimento, rg, cpf, cnh FROM 					cliente " + "WHERE id=" + fmt.Sprint(id))
	if err != nil {
		log.Fatal("Erro ao retornar cadastro", err)
	}

	for rows.Next() {
		err = rows.Scan(&cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
		if err != nil {
			log.Fatal("Erro ao pegar dados do cliente ", err)
		}
	}
	return cliente, nil
}

func Logar(email, senha string) error {
	db, err := OpenSQL()
	var log entity.Login
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	//Alterar Query para QueryRow
	rows, err := db.Query("SELECT email, senha, token FROM login " +
		"WHERE email='" + email + "' AND " + "senha='" + senha + "'")
	if err != nil {
		return fmt.Errorf("Erro ao fazer select", err)
	}
	for rows.Next() {
		err = rows.Scan(&log.Email, &log.Senha, &log.Token)
		if err != nil {
			return fmt.Errorf("Erro ao pegar dados do login", err)
		}
	}

	if log.Email == "" && log.Senha == "" && log.Token == "" {
		return fmt.Errorf("Acesso negado")
	}

	return nil
}

func GetCarrosCadastrados() ([]entity.Veiculo, error) {
	db, err := OpenSQL()
	var veiculos []entity.Veiculo
	var veiculo entity.Veiculo
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	rows, err := db.Query("SELECT  * FROM veiculo")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select para a tabela veiculos", err)
	}

	for rows.Next() {
		err = rows.Scan(&veiculo.Id, &veiculo.Modelo, &veiculo.Marca, &veiculo.Ano, &veiculo.Cor, &veiculo.Km_Litro, &veiculo.Valor_Dia, &veiculo.Valor_Hora)
		if err != nil {
			return nil, fmt.Errorf("Erro ao atribuir valores a struct veiculo")
		}
		veiculos = append(veiculos, veiculo)
	}

	if veiculos == nil {
		return nil, fmt.Errorf("Nenhum veiculo encontrado")
	}
	fmt.Println(veiculos)

	return veiculos, nil
}

func GetClientesCadastrados() ([]entity.Cliente, error) {
	var cliente entity.Cliente
	var clientes []entity.Cliente
	db, err := OpenSQL()
	if err != nil {
		return nil, fmt.Errorf("", err)
	}

	rows, err := db.Query("SELECT id, nome, sobrenome, dataNascimento, rg, cpf, cnh FROM cliente")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select na tabela cleinte: ", err)
	}

	for rows.Next() {
		err = rows.Scan(&cliente.Id, &cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
		if err != nil {
			return nil, fmt.Errorf("Erro ao passar valores para cliente")
		}
		clientes = append(clientes, cliente)
	}

	fmt.Println(clientes)
	return clientes, nil
}
