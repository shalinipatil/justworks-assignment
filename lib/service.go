package lib

import (
	"log"
	"net/http"
)

func CallAPI() *http.Response {
	url := "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Unable to reach crypto exchange rate site!!!")
	}

	if resp.Body == nil {
		log.Fatal("Body of response is empty!!!")
	}
	return resp
}
