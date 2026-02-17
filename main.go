package main

import "fmt"

func getUserInput() float64 {
	var userHeight float64
	fmt.Print("Введите данные: ")
	fmt.Scanln(&userHeight)

	return userHeight
}

func convert(
	number float64,
	currentValue string,
	targetValue string,
) {

}

func main() {
	var userInput = getUserInput()
	const usdToEur float64 = 0.8429
	const usdToRub float64 = 76.52

	const eurToRub = usdToRub / usdToEur

	fmt.Println("eurToRub", eurToRub)
}
