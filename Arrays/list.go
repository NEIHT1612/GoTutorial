package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

type TemperatureData struct {
	day1 float64
	day2 float64
	day3 float64
	day4 float64
}

// func main() {
// 	productNames := [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices)

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }

func main(){
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices[2])
}