package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func createVeiculos(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS veiculo(" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"modelo VARCHAR(50), " +
		"marca	VARCHAR(50), " +
		"ano YEAR," +
		"cor VARCHAR(20), " +
		"km_litro FLOAT, " +
		"valor_dia DOUBLE," +
		"valor_hora DOUBLE" +
		");")
	if err != nil {
		log.Fatalf("Erro ao criar tabela veiculo")
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
	rows, err := db.Query("SELECT * FROM veiculo;")
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

func InsertVeiculo(veiculo *entity.Veiculo) error {
	db, err := OpenSQL()
	if err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec("INSERT INTO veiculo(modelo, marca, ano, cor, km_litro, valor_dia, valor_hora) " +
		"VALUES ('" + veiculo.Modelo + "', '" + veiculo.Marca + "'," + veiculo.Ano + ", '" +
		veiculo.Cor + "'," + fmt.Sprintf("%f", veiculo.Km_Litro) + "," + fmt.Sprintf("%f", veiculo.Valor_Dia) + ", " +
		fmt.Sprintf("%f", veiculo.Valor_Hora) + ");")
	if err != nil {
		return fmt.Errorf("Erro ao executar insert veiculo ", err)
	}
	veiculo.Id, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Erro ao retornar id do veiculo cadastrado")
	}

	return nil
}
