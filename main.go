package main

import "fmt"

func main() {
	const usdToEur float64 = 0.8429
	const usdToRub float64 = 76.52

	const eurToRub = usdToRub / usdToEur

	fmt.Println("eurToRub", eurToRub)
}
