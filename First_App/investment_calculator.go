package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	var investmentAmount, years float64 = 1000, 10
	var expectedReturnRate float64
	const inflationRate = 6.5
	
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	
	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
