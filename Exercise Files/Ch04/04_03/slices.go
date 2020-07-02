package main

import (
	"fmt"
	"sort"
)

func main() {
	// another way of declaring and initializing arrays
	var colors = []string{"red", "green", "blue"}
	fmt.Println("here are the colors: ", colors)

	colors = append(colors, "yellow")
	fmt.Println("append item: ", colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// make() allocates & initialize, non zero storage where one can assign the value.
	num := make([]int, 5, 5)
	num[0] = 1
	num[1] = 64
	num[2] = 36
	num[3] = 55
	num[4] = 23
	fmt.Println("here is slice: ", num)

	num = append(num, 232)
	fmt.Println(num)
	fmt.Println(cap(num))
	// https://golang.org/pkg/sort/ doc ref for sort function
	sort.Ints(num)
	fmt.Println(num)
}