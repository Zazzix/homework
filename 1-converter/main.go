package main

import "fmt"

func main() {
	userCurrency, amount, exchangeCurrency := getUserInput()
	fmt.Println(userCurrency, amount, exchangeCurrency)
	result := calculateCurrency(userCurrency, amount, exchangeCurrency)
	fmt.Printf("Результат обмена: %v", result)
}

func getUserInput() (string, float64, string) {
	var userCurrency string
	fmt.Println("Введите валюту, которую хотите обменять (usd/eur/rub): ")
	fmt.Scan(&userCurrency)

	var amount float64
	fmt.Println("Введите сумму: ")
	fmt.Scan(&amount)

	var exchangeCurrency string
	fmt.Println("Введите валюту, которую хотите получить: ")
	fmt.Scan(&exchangeCurrency)

	return userCurrency, amount, exchangeCurrency

}

func calculateCurrency(userCurrency string, amount float64, exchangeCurrency string) float64 {
	var result float64
	if userCurrency == "usd" && exchangeCurrency == "rub" {
		result = amount * 97.40
	} else if userCurrency == "usd" && exchangeCurrency == "eur" {
		result = amount * 0.91
	} else if userCurrency == "eur" && exchangeCurrency == "rub" {
		result = amount * 105.45
	} else if userCurrency == "eur" && exchangeCurrency == "usd" {
		result = amount * 1.08
	} else if userCurrency == "rub" && exchangeCurrency == "usd" {
		result = amount * 0.010
	} else if userCurrency == "rub" && exchangeCurrency == "eur" {
		result = amount * 0.0095
	}
	return result
}
