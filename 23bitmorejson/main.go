package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` //we can aliases in Golang using backticks
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //omitempty --> can be used to ignore when field is null
}

func main() {
	fmt.Println("Welcome to JSON in GoLang")

	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN JS Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "ang123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	/*[
	        {
	                "Name": "React JS Bootcamp",
	                "Price": 299,
	                "Platform": "LearnCodeOnline.in",
	                "Password": "abc123",
	                "Tags": [
	                        "web-dev",
	                        "js"
	                ]
	        },
	        {
	                "Name": "MERN JS Bootcamp",
	                "Price": 199,
	                "Platform": "LearnCodeOnline.in",
	                "Password": "bcd123",
	                "Tags": [
	                        "full-stack",
	                        "js"
	                ]
	        },
	        {
	                "Name": "Angular Bootcamp",
	                "Price": 299,
	                "Platform": "LearnCodeOnline.in",
	                "Password": "ang123",
	                "Tags": null
	        }
	]
	PS D:\Users\Sandeep\Desktop\mygoLang-lco\23bitmorejson> go run .\main.go
	Welcome to JSON in GoLang
	[
	        {
	                "courseName": "React JS Bootcamp",
	                "Price": 299,
	                "website": "LearnCodeOnline.in",
	                "tags": [
	                        "web-dev",
	                        "js"
	                ]
	        },
	        {
	                "courseName": "MERN JS Bootcamp",
	                "Price": 199,
	                "website": "LearnCodeOnline.in",
	                "tags": [
	                        "full-stack",
	                        "js"
	                ]
	        },
	        {
	                "courseName": "Angular Bootcamp",
	                "Price": 299,
	                "website": "LearnCodeOnline.in"
	        }
	]*/
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "React JS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	/*JSON was valid
	main.course{Name:"React JS Bootcamp", Price:299, Platform:"LearnCodeOnline.in", Password:"", Tags:[]string{"web-dev", "js"}}*/

	//some cases where you want to add data to key value pair

	var myOnlineData map[string]interface{} //since value can be anything so use interface in Golang
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) //map[string]interface {}{"Price":299, "courseName":"React JS Bootcamp", "tags":[]interface {}{"web-dev", "js"}, "website":"LearnCodeOnline.in"}

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
		/*Key is courseName and value is React JS Bootcamp and Type is string
		Key is Price and value is 299 and Type is float64
		Key is website and value is LearnCodeOnline.in and Type is string
		Key is tags and value is [web-dev js] and Type is []interface {}*/
	}
}
