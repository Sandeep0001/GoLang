package main

import "fmt"

const LoginToken string = "asbiuhvsdb" //variable starting with capital letter, here "L" means it is publically decalred like in Java we use public

func main() {
	var username string = "Sandeep"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal) //255
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 255.45523432432
	fmt.Println(smallFloat) //255.45523
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var smallFloatPrecesion float64 = 255.45523432432
	fmt.Println(smallFloatPrecesion) //255.45523432432
	fmt.Printf("Variable is of type: %T\n", smallFloatPrecesion)

	//default values and some aliases
	var anotherVariabel int
	fmt.Println(anotherVariabel) //0
	fmt.Printf("Variable is of type: %T\n", anotherVariabel)

	var anotherString string
	fmt.Println(anotherString) //  "adds a space as default value"
	fmt.Printf("Variable is of type: %T\n", anotherString)

	//implicit type to declare variables

	var website = "learncodeonline.in"
	fmt.Println(website) //learncodeonline.in

	//no var style to declare variables
	numberOfUsers := 300000 //:= is called walrus operator and it can be used only inside methods and not global levels
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)                             //asbiuhvsdb
	fmt.Printf("Variable is of type: %T\n", LoginToken) //string

}
