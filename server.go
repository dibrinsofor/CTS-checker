package main

import (
	"flag"
	"fmt"
	"price_checker/client"
)

// I'm listening to this playlist and in GD it sounds fire.
// Magixx - Love don't cost a dime
// Vector - Early Momo. If i hear ycee next. i'll be super psyched
// SDC does not dissapoint
// Vector is so good. I think I may have him over MI
// Alpha stays undefeated. i relate so much man. i miss you so much TA

func main() {
	// var fiatcurrency, nameOfStableCoin string

	// lookup a yt of wtf flag is rn
	fiatcurrency := flag.String(
		"fiat", "NGN", "The name of the fiat I would like to pair a coin up against",
	)

	nameOfStableCoin := flag.String(
		"coin", "NGNT", "The name of the coin I would like to know the price of",
	)

	// fiat := &fiatcurrency
	// token := &nameOfStableCoin
	// it prints out memory addresses
	// 	0xc000044270
	// 0xc000044280

	flag.Parse()

	// 	price_checker/client.FetchPrice({0xc000052060, 0xc000044200}, {0x4, 0xb47973})
	//         C:/Users/Dibri Nsofor/Desktop/Go/src/price_checker/client/pcclient.go:28 +0x1da
	// main.main()
	//         C:/Users/Dibri Nsofor/Desktop/Go/src/price_checker/server.go:38 +0xf4
	// exit status 2
	// fmt.Println(*fiatcurrency)
	// fmt.Println(*nameOfStableCoin)
	// e be here the problem dey come from
	pegged_coin, _ := client.FetchPrice(*fiatcurrency, *nameOfStableCoin)

	fmt.Println(pegged_coin)
}
