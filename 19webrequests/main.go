package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO webrequest in GoLang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response) //Response is of type *http.Response

	defer response.Body.Close() //its caller's responsibilty to always close the connection afer use

	databytes, err := ioutil.ReadAll(response.Body)

	content := string(databytes)

	fmt.Println(content)

}
