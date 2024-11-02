package main

import (
	"fmt"
	"strconv"
)

func main() {
	currencies := map[string]map[string]float64{
		"EUR": {
			"USD": 1.08,
			"RUB": 105.45,
		},
		"USD": {
			"EUR": 0.91,
			"RUB": 97.40,
		},
		"RUB": {
			"EUR": 0.0095,
			"USD": 0.010,
		},
	}
	userCurrency, amount, exchangeCurrency := getUserInput()
	fmt.Println(userCurrency, amount, exchangeCurrency)
	result := calculateCurrency(userCurrency, amount, exchangeCurrency, currencies)
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

func calculateCurrency(userCurrency string, amount float64, exchangeCurrency string, currencies map[string]map[string]float64) float64 {
	var result float64
	if userCurrency == "usd" && exchangeCurrency == "rub" {
		result = amount * currencies["USD"]["RUB"]
	} else if userCurrency == "usd" && exchangeCurrency == "eur" {
		result = amount * currencies["USD"]["EUR"]
	} else if userCurrency == "eur" && exchangeCurrency == "rub" {
		result = amount * currencies["EUR"]["RUB"]
	} else if userCurrency == "eur" && exchangeCurrency == "usd" {
		result = amount * currencies["EUR"]["USD"]
	} else if userCurrency == "rub" && exchangeCurrency == "usd" {
		result = amount * currencies["RUB"]["USD"]
	} else if userCurrency == "rub" && exchangeCurrency == "eur" {
		result = amount * currencies["RUB"]["EUR"]
	}
	return result
}
