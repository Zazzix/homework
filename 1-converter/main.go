package main

import (
	"fmt"
	"strconv"
)

func main() {
	userCurrency, amount, exchangeCurrency := getUserInput()
	fmt.Println(userCurrency, amount, exchangeCurrency)
	result := calculateCurrency(userCurrency, amount, exchangeCurrency)
	fmt.Printf("Результат обмена: %v\n", result)
}

func getUserInput() (string, float64, string) {
	var userCurrency string
	fmt.Print("Введите валюту, которую хотите обменять (usd/eur/rub): ")
	fmt.Scan(&userCurrency)
	for {
		if userCurrency == "usd" || userCurrency == "eur" || userCurrency == "rub" {
			break
		} else {
			fmt.Print("Введите валюту, которую хотите обменять (usd/eur/rub): ")
			fmt.Scan(&userCurrency)
		}
	}

	var amount string
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	var userAmount float64
	for {
		userAmount, _ = strconv.ParseFloat(amount, 64)
		if userAmount > 0 {
			break
		} else {
			fmt.Print("Введите сумму: ")
			fmt.Scan(&amount)
		}
	}

	var exchangeCurrency string
	switch userCurrency {
	case "usd":
		fmt.Print("Введите валюту, которую хотите получить (eur/rub): ")
		fmt.Scan(&exchangeCurrency)
	case "eur":
		fmt.Print("Введите валюту, которую хотите получить (usd/rub): ")
		fmt.Scan(&exchangeCurrency)
	case "rub":
		fmt.Print("Введите валюту, которую хотите получить (usd/eur): ")
		fmt.Scan(&exchangeCurrency)
	}

	for {
		if exchangeCurrency == userCurrency {
			fmt.Printf("Вы не можете обменять %v на %v. Выберите другую валюту: ", userCurrency, exchangeCurrency)
			fmt.Scan(&exchangeCurrency)
		} else {
			if exchangeCurrency == "usd" || exchangeCurrency == "eur" || exchangeCurrency == "rub" {
				break
			} else {
				fmt.Print("Введите валюту, которую хотите получить: ")
				fmt.Scan(&exchangeCurrency)
			}
		}
	}

	return userCurrency, userAmount, exchangeCurrency

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
