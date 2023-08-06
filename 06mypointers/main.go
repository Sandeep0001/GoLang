package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int // is used to define pointers
	// fmt.Println("Value of pointer is ", ptr) //Value of pointer is  <nil>

	myNumber := 23

	var ptr = &myNumber                             //& is used as reference of memory for defined variable
	fmt.Println("Value of actual pointer is", ptr)  //Value of actual pointer is 0xc000016098
	fmt.Println("Value of actual pointer is", *ptr) //Value of actual pointer is 23

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber) //New value is 46
}
