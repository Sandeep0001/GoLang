package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to learning on Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of  fruitList is %T\n", fruitList) //Type of  fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) //[Apple Tomato Peach Mango Banana]

	fruitList1 := append(fruitList[1:])
	fmt.Println(fruitList1) // [Tomato Peach Mango Banana]

	fruitList2 := append(fruitList[1:3])
	fmt.Println(fruitList2) //[Peach Mango] --> last index range is not inclusive hencce slice gives us value of 1,2

	fruitList3 := append(fruitList[:3])
	fmt.Println(fruitList3) //[Apple Tomato Peach]

	//allocating memory using make() to slice
	highScores := make([]int, 4) //allocating memory for slice

	highScores[0] = 244
	highScores[1] = 245
	highScores[2] = 246
	highScores[3] = 230
	//highScores[4] = 344  //will throw index out of range error

	highScores = append(highScores, 988, 133, 234) // using append method we can allocate extra memory in slice

	fmt.Println(highScores) //[244 245 246 230 988 133 234]

	sort.Ints(highScores)
	fmt.Println(highScores)                     //[133 230 234 244 245 246 988]
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	//how to remove a value from slices based on index
	var courses = []string{"reactJS", "javascript", "python", "ruby", "swift"}
	fmt.Println(courses) //[reactJS javascript python ruby swift]

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[reactJS javascript ruby swift]
}
