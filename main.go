package main

import (
	"fmt"
	"justworks_assignment/lib"
	"log"
	"os"
	"strconv"
)

const currency string = "$"

func main() {
	arguments := os.Args

	amount, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil || amount <= 0 {
		log.Fatal("Amount is not valid!!!")
	}
	crypto1 := arguments[2]
	crypto2 := arguments[3]

	resp := lib.CallAPI()

	obj := lib.RespUnmarshaler(resp)

	per70OfAmount, per30ofAmount, crypto1Buy, crypto2Buy := lib.MoneyExchanger(arguments[1], obj, crypto1, crypto2)
	fmt.Println(fmt.Sprint(currency, per70OfAmount), "=>", crypto1Buy, crypto1)
	fmt.Println(fmt.Sprint(currency, per30ofAmount), "=>", crypto2Buy, crypto2)

}
