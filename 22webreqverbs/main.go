package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get request in Golang")

	PerformGetRequest()
	PerfromPostJsonRequest()
	PerfromPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode) //200
	fmt.Println("Content le length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content)) //{"message":"Hello from lco.in"}

	//2nd approach
	var responseString strings.Builder
	content1, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content1)

	fmt.Println("Byte count is:", byteCount) //43 --> which is content length
	fmt.Println(responseString.String())     //{"message":"Hello from lco.in"}
}

func PerfromPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName" : "Let's go with Golang",
			"price" : 0,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerfromPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "Sandeep")
	data.Add("lasttName", "Shrinivas")
	data.Add("email", "Sandeep@go.dev.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
