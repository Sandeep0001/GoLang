package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime) //2023-07-12 15:25:03.1377288 +0530 IST m=+0.001067901

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //07-12-2023 15:33:17 Wednesday

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)                             //2020-08-10 23:23:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 Monday")) //08-12-2020 Wednesday
}
