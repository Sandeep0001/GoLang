package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in GoLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist is", fruitList)             //Fruitlist is [Apple Tomato  Peach] //index 2 has blank space
	fmt.Println("Fruitlist length is", len(fruitList)) //Fruitlist length is 4 --> it always gives the initialized length

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is", vegList)             //Veg list is [potato beans mushroom]
	fmt.Println("Veg list length is", len(vegList)) //Veg list length is 5
}
