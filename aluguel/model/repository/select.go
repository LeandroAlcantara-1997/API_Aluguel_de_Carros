package repository

import (
	"fmt"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func GetByIdCliente(id int) (entity.Cliente, error) {
	var cliente entity.Cliente

	db, err := OpenSQL()
	if err != nil {
		return cliente, fmt.Errorf("%v", err)
	}

	rows := db.QueryRow("SELECT nome, sobrenome, dataNascimento, rg, cpf, cnh FROM 	cliente " + "WHERE id= " + fmt.Sprint(id))

	err = rows.Scan(&cliente.Nome, &cliente.Sobrenome, &cliente.Data_Nascimento, &cliente.RG, &cliente.CPF, &cliente.CNH)
	if err != nil {
		return cliente, fmt.Errorf("Erro ao pegar dados do cliente ", err)
	}

	return cliente, nil
}

func Logar(email, senha string) error {
	var log entity.Login

	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	token, err := entity.GeraToken(email + senha)
	if err != nil {
		return fmt.Errorf("%#v", err)
	}
	rows := db.QueryRow("SELECT token FROM login " +
		"WHERE token='" + token + "'")

	err = rows.Scan(&log.Token)
	if err != nil {
		return fmt.Errorf("Acesso negado", err)
	}

	return nil
}

func GetEmailToSenha(email string) error {
	db, err := OpenSQL()
	var log entity.Login
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	rows := db.QueryRow("SELECT email FROM login " +
		"WHERE email='" + email + "'")
	fmt.Print(rows)
	err = rows.Scan(&log.Email)
	if err != nil {
		return fmt.Errorf("Email não cadastrado", err)
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

func GetAlugueis() ([]entity.Aluguel, error){
	db, err := OpenSQL()
	var alugueis []entity.Aluguel
	var aluguel entity.Aluguel
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	rows, err := db.Query("SELECT  * FROM aluguel")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select para a tabela alugueis", err)
	}

	for rows.Next() {
		err = rows.Scan(&aluguel.Data_Inicio, &aluguel.Data_Retorno, &aluguel.Valor_Total)
		if err != nil {
			return nil, fmt.Errorf("Erro ao atribuir valores a struct aluguel")
		}
		alugueis = append(alugueis, aluguel)
	}

	if alugueis == nil {
		return nil, fmt.Errorf("Nenhum veiculo encontrado")
	}
	fmt.Println(alugueis)

	return alugueis, nil
}


func LogarAdmin(admin *entity.Admin) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("Erro open sql", err)
	}

	fmt.Println(admin.Token)

	rows := db.QueryRow("SELECT token FROM admin " +
		"WHERE token = '" + admin.Token + "'")

	err = rows.Scan(&admin.Token)
	if err != nil {
		return fmt.Errorf("Acesso negado", err)
	}
	return nil
}