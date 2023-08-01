package main

import (
	"github.com/mayank-sheoran/zerodha-sdk-go/pkg"
	"github.com/zerodha/gokiteconnect/v4/models"
)

func main() {
	tokens := []string{"NSE:NIFTY BANK", "NSE:INFY"}
	kiteConnect := pkg.KiteConnect("")
	quote, err := kiteConnect.GetQuote(tokens...)
	println(quote["NSE:NIFTY BANK"].SellQuantity)
	println(quote["NSE:INFY"].SellQuantity)
	if err != nil {
		panic(err.Error())
	}
	marketDepth := map[string]models.Depth{}
	for _, token := range tokens {
		marketDepth[token] = quote[token].Depth
	}
}
