package main

import (
	"fmt"
	"strings"
)

/*
Сделать меню с шагами
- ввод исходно валюты (подсказывать варианты, если ошибка, заново вводит
- ввод числа, если ошибка, заново вводит
- ввод целевой валюты, если ошибка, заново вводит
*/

func getUserStringInput(valuteText string) string {
	var inputValue string
	var inputText string = fmt.Sprintf("Введите %s валюту (rub / eur / usd): ", valuteText)

	for {
		fmt.Printf(inputText)

		var _, err = fmt.Scanln(&inputValue)
		inputValue = strings.TrimSpace(inputValue)
		fmt.Println("err", err)

		if err != nil || inputValue != "rub" && inputValue != "eur" && inputValue != "usd" {
			fmt.Println("Введены некорректные данные")

			continue
		}

		return inputValue
	}
}

func getUserNumberInput() float64 {
	var inputValue float64
	var inputText string = "Введите число с точкой: "

	for {
		fmt.Printf(inputText)
		var _, err = fmt.Scanln(&inputValue)

		if err != nil {
			var dummy string
			fmt.Scanln(&dummy)

			fmt.Println("Введены некорректные данные")

			continue
		}

		return inputValue
	}
}

func convert(
	number float64,
	currentValue string,
	targetValue string,
) float64 {
	const usdToEur float64 = 0.85
	const usdToRub float64 = 76.65
	const eurToRub = usdToRub / usdToEur // 90.18
	var converterValue float64

	if currentValue == targetValue {
		return number
	}

	switch {
	case currentValue == "rub":
		if targetValue == "eur" {
			converterValue = usdToEur / usdToRub
		}

		if targetValue == "usd" {
			converterValue = usdToEur / usdToRub
		}

	case currentValue == "usd":
		if targetValue == "eur" {
			converterValue = usdToEur
		}

		if targetValue == "rub" {
			converterValue = usdToRub
		}

	case currentValue == "eur":
		if targetValue == "usd" {
			converterValue = eurToRub / usdToRub
		}

		if targetValue == "rub" {
			converterValue = eurToRub
		}
	}

	return converterValue * number
}

func main() {
	var numberValue = getUserNumberInput()
	fmt.Println("numberValue", numberValue)

	var currentValue = getUserStringInput("исходную")
	fmt.Println("currentValue", currentValue)

	var targetValue = getUserStringInput("целевую")
	fmt.Println("targetValue", targetValue)

	var convertValue = convert(numberValue, currentValue, targetValue)

	fmt.Println("convertValue", convertValue)
}
