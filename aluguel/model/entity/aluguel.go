package entity


type Aluguel struct {
	Id_Cliente		int
	Veiculo			Veiculo
	Data_Inicio	string
	Data_Retorno	string
	Valor_Total		float64
}


