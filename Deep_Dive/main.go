package main

import (
	"fmt"
)

type transformN func(int) int
type anotherN func(int, []string, map[string] string) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}
	dNumbers := transformNumbers(&numbers, double)
	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println(dNumbers)
	fmt.Println(tNumbers)

	transformN1 := getTransformerFunction(&numbers)
	transformN2 := getTransformerFunction(&moreNumbers)
	transformeNumbers1 := transformNumbers(&numbers, transformN1)
	transformeNumbers2 := transformNumbers(&moreNumbers, transformN2)
	fmt.Println(transformeNumbers1)
	fmt.Println(transformeNumbers2)

	fmt.Println(createTransformer(3)(4))
}

func transformNumbers(numbers *[]int, transform transformN) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}
func double(number int) int{
	return number * 2
}
func triple(number int) int{
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformN{
	if (*numbers)[0] == 1{
		return double
	}else{
		return triple
	}
}

func createTransformer(factor int) func(int) int{
	return func(number int) int{
		return number * factor
	}
}