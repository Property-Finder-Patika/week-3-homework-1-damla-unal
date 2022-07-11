package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	USD
)

// you're going to display currency exchange ratios against TL
func main() {
	//currencyRates that contains the conversion ratios
	currencyRates := [...]float64{
		EUR: 17.45,
		USD: 17.38,
	}

	// get TL amount to be converted from command line
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter the amount to be converted : ")
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f TL is %.2f EUR\n", amount, amount/currencyRates[EUR])
	fmt.Printf("%.2f TL is %.2f USD\n", amount, amount/currencyRates[USD])
}
