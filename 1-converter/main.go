package main

import "fmt"

func main() {
	const USDtoEUR = 0.91 * 1
	const USDtoRUB = 96.20 * 1
	const EURtoRUB = USDtoEUR * USDtoRUB
	getUserInput()
}

func getUserInput() {
	var currency string
	fmt.Scan(&currency)
	fmt.Println(currency)
}

func calculateCurrency(amount float64, currency1 string, currency2 string) float64 {

}
