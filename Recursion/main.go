package main

import "fmt"

func main() {
	factory := factorial(5)
	fmt.Println(factory)

	numbers := []int{5, 6, 7, 8}
	sum := sumup(1, 2, 3, 4)
	newSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(newSum)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func sumup(numbers ...int) int{
	sum := 0
	for _, val := range numbers{
		sum += val
	}
	return sum
}