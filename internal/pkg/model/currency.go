package model

type CurrencyRes struct {
	Data    map[string]float64 `json:"data"`
	Service string             `json:"service"`
}
