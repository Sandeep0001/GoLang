package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Functions in Golang")

	result := adder(3, 5)
	fmt.Println("Rsult is:", result) //8

	fmt.Println("Proadder result is:", proAdder(2, 3, 4, 5)) //14

	mulResult, proMulMessage := proMultiplierWithMessage(2, 3, 4)
	fmt.Println(mulResult, ",", proMulMessage) //24 , Hi this is pro multiplication
}

func greeter() {
	fmt.Println("Namaste!")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int { // ... is referred to Variadic function where trailing arguments can be passed
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func proMultiplierWithMessage(values ...int) (int, string) {
	total := 1

	for _, val := range values {
		total *= val
	}

	return total, "Hi this is pro multiplication"
}
