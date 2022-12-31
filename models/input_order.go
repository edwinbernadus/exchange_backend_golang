package models

type InputOrder struct {
	RequestRate  float64
	CurrencyPair string
	//Uuid         string
	Amount   float64
	IsBuy    bool
	UserName string
}
