package main

import (
	"fmt"
)

func main() {
	// ordered collection of values
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Yellow"
	fmt.Println("here is the arr: ",colors)
	fmt.Println(colors[1])

	// single line decleration of an array
	var numbers = [5]int{5,4,3,2,1}
	fmt.Println("the numbers are: ",numbers)

	fmt.Println("no. of colors : ", len(colors))
	fmt.Println("no. of numbers : ", len(numbers))
}