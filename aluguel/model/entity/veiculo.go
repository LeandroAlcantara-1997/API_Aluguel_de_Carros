package entity

type Veiculo struct {
	Id			int `json:"id"`
	Modelo		string `json:"modelo"`
	Marca		string
	Ano			string
	Cor			string
	Km_Litro 	float64
	Valor_Dia	float64
	Valor_Hora	float64
}