package lib

import (
	"log"

	"github.com/shopspring/decimal"
)

const split70val float64 = 0.7
const split30val float64 = 0.3

func MoneyExchanger(obj *CryptoResponse, arg string, crypto1 string, crypto2 string) (per70OfAmount, per30ofAmount, crypto1Buy, crypto2Buy string) {
	if arg == "" || crypto1 == "" || crypto2 == "" || obj == nil {
		log.Fatal("Exchanger can not proceed with empty value!!!")
	}

	amount, err := decimal.NewFromString(arg)
	if err != nil || amount.IsZero() || amount.IsNegative() {
		log.Fatal("Amount is not valid")
	}

	decimalOf70Per := amount.Mul(decimal.NewFromFloat(split70val))
	decimalOf30Per := amount.Mul(decimal.NewFromFloat(split30val))

	var exchangeR1, exchangeR2, crypto1Amount, crypto2Amount decimal.Decimal
	for key, val := range obj.Data.Rates {
		if key == crypto1 {
			exchangeR1, err = decimal.NewFromString(val)
			if err != nil {
				log.Fatal("Unable to parse data for crypto1")
			}
			crypto1Amount = exchangeR1.Mul(decimalOf70Per)
		}
		if key == crypto2 {
			exchangeR2, err = decimal.NewFromString(val)
			if err != nil {
				log.Fatal("Unable to parse data for crypto2")
			}
			crypto2Amount = exchangeR2.Mul(decimalOf30Per)
		}
	}

	validateExchangeVals(decimalOf70Per, decimalOf30Per, crypto1Amount, crypto2Amount)

	return decimalOf70Per.Round(2).String(), decimalOf30Per.Round(2).String(), crypto1Amount.Round(5).String(), crypto2Amount.Round(5).String()

}

func validateExchangeVals(decimalOf70Per, decimalOf30Per, crypto1Amount, crypto2Amount decimal.Decimal) {
	if decimalOf70Per.IsZero() || decimalOf30Per.IsZero() || crypto1Amount.IsZero() || crypto2Amount.IsZero() {
		log.Fatal("Exchange return zero value!!!")
	}
	if decimalOf70Per.IsNegative() || decimalOf30Per.IsNegative() || crypto1Amount.IsNegative() || crypto2Amount.IsNegative() {
		log.Fatal("Exchange return negative value!!!")
	}
}
