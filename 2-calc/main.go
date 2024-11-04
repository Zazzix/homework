package main

import (
	"fmt"
	"strconv"
	"slices"
	"strings"
)

func main() {
	operation, numbers := getUserInput()
	fmt.Println(operation, numbers)
	result := calculate(operation, numbers)
	fmt.Println(result)
}

func getUserInput() (string, []float64) {
	var operation string
	var input string
	fmt.Print("Выберите тип операции (avg/sum/med): ")
	fmt.Scan(&operation)
	operation = strings.ToLower(operation)
	for {
		if operation == "sum" || operation == "avg" || operation == "med" {
			break
		} else {
			fmt.Print("Выберите тип операции (avg/sum/med): ")
			fmt.Scan(&operation)
		}
	}
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&input)
	parsedString := strings.Split(input, ",")
	numbers := make([]float64, 0, len(parsedString))

	for _, element := range parsedString {
		element, _ := strconv.ParseFloat(element, 64)
		numbers = append(numbers, element)
	}
	slices.Sort(numbers)
	return operation, numbers
}

func calculate(operation string, numbers []float64) float64 {
	var result float64
	arrLength := len(numbers)
	switch operation {
	case "sum":
		for _, value := range numbers {
			result += value
		}
	case "avg":
		for _, value := range numbers {
			result += value
		}
		result = result / float64(len(numbers))
	case "med":
		if arrLength%2 == 0 {
			result = (numbers[arrLength/2-1] + numbers[arrLength/2]) / 2
		} else {
			result = numbers[arrLength/2]
		}
	}
	return result
}
