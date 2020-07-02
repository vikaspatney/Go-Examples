package main

import "fmt"

func main() {
	var ptr *int

	if ptr != nil {
		fmt.Println("value of ptr:", *ptr)
	} else {
		fmt.Println("ptr is nil")
	}
	// assigning the value v to ptr
	var v int = 43
	ptr = &v
	if ptr != nil {
		fmt.Println("value of ptr:", *ptr)
	} else {
		fmt.Println("p is nil")
	}

	var val1 float64 = 34.40
	// implicit decleration of the pointer
	ptr1 := &val1
	fmt.Println("val1 is :", *ptr1)

	// modifying the value using the pointer 
	*ptr1 = *ptr1 / 30
	fmt.Println("the value after pointer modifying :", *ptr1)
	// now checking if the original val of val1 is also modified
	fmt.Println("the val1 =", val1)

}
