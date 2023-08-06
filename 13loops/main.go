package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days) //[Sunday Tuesday Wednesday Friday Saturday]

	//1.
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
		/*Sunday
		Tuesday
		Wednesday
		Friiday
		Saturday*/
	}

	//2
	for i := range days {
		fmt.Println(days[i]) //here i returns the index value
		/*Sunday
		Tuesday
		Wednesday
		Friday
		Saturday*/
	}

	//3
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
		/*index is 0 and value is Sunday
		index is 1 and value is Tuesday
		index is 2 and value is Wednesday
		index is 3 and value is Friday
		index is 4 and value is Saturday*/
	}

	//4
	for _, day := range days {
		fmt.Println("value is", day)
		/*value is Sunday
		value is Tuesday
		value is Wednesday
		value is Friday
		value is Saturday*/
	}

	//like while loop
	//pre decrement / increment operators are not allowed in Golang
	rougeVlaue := 1

	for rougeVlaue < 10 {
		fmt.Println("Value is:", rougeVlaue)
		rougeVlaue++
		/*Value is: 1
		Value is: 2
		Value is: 3
		Value is: 4
		Value is: 5
		Value is: 6
		Value is: 7
		Value is: 8
		Value is: 9*/
	}

	//break
	val1 := 1
	for val1 < 10 {

		if val1 == 5 {
			break
		}

		fmt.Println("Value is:", val1)
		val1++
		/*Value is: 1
		Value is: 2
		Value is: 3
		Value is: 4*/
	}

	//continue
	val2 := 1
	for val2 < 10 {
		if val2 == 5 {
			val2++
			continue
		}

		fmt.Println("Value is:", val2)
		val2++
		/*Value is: 1
		Value is: 2
		Value is: 3
		Value is: 4
		Value is: 6
		Value is: 7
		Value is: 8
		Value is: 9*/
	}

	//goto
	val3 := 1
	for val3 < 10 {

		if val3 == 2 {
			goto abc
		}

		fmt.Println("Value is:", val3)
		val3++
		/*Value is: 1
		This is goto statement*/
	}

abc:
	fmt.Println("This is goto statement")

}
