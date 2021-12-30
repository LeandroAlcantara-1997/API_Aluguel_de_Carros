package entity

type Veiculo struct {
	Id         int
	Modelo     string  `json:"modelo"`
	Marca      string  `json:"marca"`
	Ano        string  `json:"ano"`
	Cor        string  `json:"cor"`
	Km_Litro   float64 `json:"km_litro"`
	Valor_Dia  float64 `json: "valor_dia"`
	Valor_Hora float64 `json: "valor_hora"`
}
