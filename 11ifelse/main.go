package main

import "fmt"

func main() {
	fmt.Println("if else in Golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login counts"
	}

	fmt.Println(result) //watch out

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	} //Number is odd

	//assigning the value and checking the condition syntax
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	} //Number is less than 10
}
