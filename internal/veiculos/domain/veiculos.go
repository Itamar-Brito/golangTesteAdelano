package domain

//struct para veiculos

type Veiculos struct {
	ID     int    `json:"id" binding:"required"`
	Placa  string `json:"placa" binding:"required"`
	Modelo string `json:"modelo" binding:"required"`
	Ano    string `json:"ano" binding:"required"`
}
