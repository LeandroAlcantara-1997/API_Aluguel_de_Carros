package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/entity"
)

func createAluguel(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS aluguel(" +
		"data_inicio DATETIME, " +
		"data_retorno DATETIME, " +
		"total DOUBLE, " +
		"fk_veiculo INT, " +
		"fk_cliente INT" +
		");")

	if err != nil {
		log.Fatalf("Erro ao criar tabela aluguel %#v", err)
	}

	return nil
}

func GetAlugueis() ([]entity.Aluguel, error) {
	db, err := OpenSQL()
	var alugueis []entity.Aluguel
	var aluguel entity.Aluguel
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	rows, err := db.Query("SELECT * FROM aluguel")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select para a tabela alugueis %#v", err)
	}

	for rows.Next() {
		err = rows.Scan(&aluguel.Inicio, &aluguel.Retorno, &aluguel.Valor_Total)
		if err != nil {
			return nil, fmt.Errorf("Erro ao atribuir valores a struct aluguel")
		}
		alugueis = append(alugueis, aluguel)
	}

	if alugueis == nil {
		return nil, fmt.Errorf("Nenhum veiculo aluguel encontrado")
	}

	return alugueis, nil
}


func InsertAluguel(aluguel entity.Aluguel) error {
	db, err := OpenSQL()
	if err != nil {
		return err
	}
	cliente := fmt.Sprintf("%d", aluguel.Id_Cliente)
	veiculo := fmt.Sprintf("%d", aluguel.Id_Veiculo)
	_, err = db.Exec("INSERT INTO aluguel (fk_cliente, fk_veiculo, data_inicio, data_retorno, total) " +
		"VALUES (" + cliente + ", " + veiculo + ", '" + aluguel.Inicio + "', '" + aluguel.Retorno + "', " + fmt.Sprintf("%.2f", aluguel.Valor_Total) + ");")

	if err != nil {
		return fmt.Errorf("Erro ao realizar insert na tabela aluguel: %#v", err)
	}
	return nil
}

func DeleteAluguel(id string) error {
	db, err := OpenSQL()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM aluguel " +
		"WHERE fk_cliente = '" + id + "';")
	if err != nil {
		return fmt.Errorf("Erro ao deletar alugueis %#v", err)
	}

	return nil
}

func GetAlugadosCliente(id string) ([]entity.Veiculo, error) {
	var veiculo entity.Veiculo
	var veiculos []entity.Veiculo
	db, err := OpenSQL()
	if err != nil {
		return veiculos, err
	}
	rows, err := db.Query("SELECT id, modelo, marca, ano, cor, km_litro, valor_dia, valor_hora FROM veiculo " +
		"INNER JOIN aluguel " +
		"ON veiculo.id=aluguel.fk_veiculo AND aluguel.fk_cliente=" + id + ";")

	for rows.Next() {
		err = rows.Scan(&veiculo.Id, &veiculo.Modelo, &veiculo.Marca,
			&veiculo.Ano, &veiculo.Cor, &veiculo.Km_Litro, &veiculo.Valor_Dia, &veiculo.Valor_Hora)
		if err != nil {
			return veiculos, fmt.Errorf("Erro ao pegar dados do cliente")
		}
		veiculos = append(veiculos, veiculo)
	}

	if veiculos == nil {
		return veiculos, fmt.Errorf("Nenhum veiculo encotrado")
	}

	return veiculos, nil
}
