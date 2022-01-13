package model

import (
	"fmt"
)

type CurrencyResponse struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	High   string `json:"high"`
	Supply string `json:"circulating_supply"`
}

func (r CurrencyResponse) CLIresponse() string {
	res := fmt.Sprintf(
		"Name: %s\nPrice: %s\nRank: %s\nHigh: %s\nSupply: %s\n",
		r.Name, r.Price, r.Rank, r.High, r.Supply)

	return res
}
