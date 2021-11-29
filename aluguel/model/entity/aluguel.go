package entity


type Aluguel struct {
	Id_Cliente		int
	Veiculo			Veiculo
	Data_Aluguel	string
	Data_Retorno	string
	Valor_Toal		float64
}


