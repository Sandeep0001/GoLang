package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages) //List of all languages: map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS short for:", languages["JS"])    //JS short for: Javascript

	delete(languages, "RB")
	fmt.Println(languages) //map[JS:Javascript PY:Python]

	//loops are intresting golang
	//for each loop
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value) /*For key JS, value is Javascript
		For key PY, value is Python*/
	}

	//using comma ok syntax
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value) /*For key v, value is Python
		For key v, value is Javascript*/
	}
}
