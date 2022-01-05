package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createAluguel(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS aluguel(" +
		"data_inicio TIMESTAMP, " +
		"data_retorno TIMESTAMP, " +
		"total DOUBLE, " +
		"fk_veiculo INT, " +
		"fk_cliente INT" +
		");")

	if err != nil {
		log.Fatalf("Erro ao criar tabela aluguel ", err)
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
		return nil, fmt.Errorf("Nenhum veiculo aluguel encontrado")
	}
	fmt.Println(alugueis)

	return alugueis, nil
}

func GetAlugueisCliente() ([]entity.Veiculo, error) {
	var veiculo entity.Veiculo
	var veiculos []entity.Veiculo
	db, err := OpenSQL()
	if err != nil {
		return nil, fmt.Errorf("%#v", err)
	}

	rows, err := db.Query("SELECT * FROM veiculo " +
		"INNER JOIN aluguel " +
		"ON veiculo.id = aluguel.fk_veiculo;")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar select para a tabela alugueis", err)
	}

	for rows.Next() {
		err = rows.Scan(&veiculo.Modelo, &veiculo.Marca, &veiculo.Ano, &veiculo.Cor, &veiculo.Id, &veiculo.Valor_Dia, &veiculo.Valor_Hora)
		if err != nil {
			return nil, fmt.Errorf("Erro ao atribuir valores a struct veiculo")
		}
		veiculos = append(veiculos, veiculo)
	}

	if veiculos == nil {
		return nil, fmt.Errorf("Nenhum veiculo aluguel encontrado")
	}
	fmt.Println(veiculos)

	return veiculos, nil
}
