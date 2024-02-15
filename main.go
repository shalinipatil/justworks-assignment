package main

import (
	"fmt"
	"justworks_assignment/lib"
	"log"
	"os"
)

const currency string = "$"

func main() {
	arguments := os.Args

	if len(arguments) <= 3 {
		log.Fatal("Arguments not passed properly.")
	}
	crypto1 := arguments[2]
	crypto2 := arguments[3]

	resp := lib.CallAPI()

	obj := lib.RespUnmarshaler(resp)

	per70OfAmount, per30ofAmount, crypto1Buy, crypto2Buy := lib.MoneyExchanger(obj, arguments[1], crypto1, crypto2)
	fmt.Println(fmt.Sprint(currency, per70OfAmount), "=>", crypto1Buy, crypto1)
	fmt.Println(fmt.Sprint(currency, per30ofAmount), "=>", crypto2Buy, crypto2)

}
