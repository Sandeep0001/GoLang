package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GoLang")
	content := "This needs to go in file - LearnCodeOnline.in"

	file, err := os.Create("./myLcoGoFile.txt")

	if err != nil {
		panic(err) //panic will shutdown the program and displays error
	}

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("length is: ", length)

	defer file.Close()
	readFile("./myLcoGoFile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
