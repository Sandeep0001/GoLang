package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//structs are same as class in Java
	//no inheritance in Golang
	//no Super or parent

	sandeep := User{"Sandeep", "san@abc.com", true, 20}
	fmt.Println(sandeep)                                                    //{Sandeep san@abc.com true 20}
	fmt.Printf("Sandeep details are: %+v\n", sandeep)                       //Sandeep details are: {Name:Sandeep Email:san@abc.com status:true Age:20}
	fmt.Printf("Name is %v and email is %v\n", sandeep.Name, sandeep.Email) //Name is Sandeep and email is san@abc.com
	sandeep.GetStatus()                                                     //Is user active: true                                              //
	sandeep.NewMail()                                                       //Email of this user is: test123@avc.com
	fmt.Printf("Name is %v and email is %v\n", sandeep.Name, sandeep.Email) //Name is Sandeep and email is san@abc.com

}

//how to defne a struct in Golang --> use keyword type
type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

//defining the methods
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)
}

func (u User) NewMail() {
	u.Email = "test123@avc.com"
	fmt.Println("Email of this user is:", u.Email)
}
