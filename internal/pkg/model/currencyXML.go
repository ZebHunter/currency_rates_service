package model

type CBRes struct {
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	CharCode string `xml:"CharCode"`
	Exchange string `xml:"Value"`
}
