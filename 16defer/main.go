package main

import "fmt"

func main() {
	defer fmt.Println("Language") //with defer keyword this statement will be executed last from the function, it works in LIFO structure
	fmt.Println("Defer in Golang")
	/*Defer in Golang
	Language*/
	//*****************************************//
	printDefers()
	/*Defer in Golang
	one
	two
	three
	Language*/

	//*****************************************//
	numberDefer()
	/*Defer in Golang
	one
	two
	three
	4
	3
	2
	1
	0
	Language*/

}

func printDefers() {
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
}

func numberDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
