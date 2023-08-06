package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //using wait group the execution time is faster
var mut sync.Mutex    //Mutex serves the purpose of lock and unlock the thread

func main() {
	go greeter("Hello") //go keyword is used to specify goroutine
	greeter("World")
	/*Hello
	World
	World
	Hello
	Hello
	World
	World
	Hello
	Hello
	World
	World*/

	//using wait groups
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) //waitgroup has methods for Add, wait and done.
	}
	wg.Wait() //asks main method to wait until all threads are finished
	fmt.Println(signals)
	/*200 status code for https://go.dev
	200 status code for https://lco.dev
	200 status code for https://fb.com
	200 status code for https://google.com
	[test https://go.dev https://lco.dev https://fb.com https://google.com]*/

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done() //this statement will be exeuted in the end and will tell wait group that exe of a thread is completed

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}
}
