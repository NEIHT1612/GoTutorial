package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//   => Show error message & exit if invalid input is provided
//   - No negative numbers
//   - Not 0
// 2) Store calculated results into file

func main() {

	revenue, err := getUserInput("Revenue:")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses:")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate:")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
	storeResults(ebt, profit, ratio)

	//Output String
	// formattedEBT := fmt.Sprintf("FormattedEBT: %.1f\n", ebt)
	// formattedPRF := fmt.Sprintf("FormattedPRF: %.1f\n", profit)
	// fmt.Print(formattedEBT, formattedPRF)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if(userInput <= 0){
		return 0, errors.New("Value must be a positive number.")
	}
	return userInput, nil
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
