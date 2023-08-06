package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "Welcome to user input"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our Pizza:")

	//comma ok / error ok syntax
	input, _ := reader.ReadString('\n') //we can use input, err := which is similar to try catch in Java
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of this rating is %T", input) //string
}
