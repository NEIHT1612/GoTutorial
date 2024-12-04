package main

import "fmt"

func main() {
	age := 10

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	getLoliYears(agePointer)
	fmt.Println(age)
}

func getLoliYears(age *int) {
	*age = *age - 1
}