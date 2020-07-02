package main

import "fmt"

func main() {

	var x float64 = 42
	var result string
	
	if x < 0 {
		result = "Less than zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println("Result: ", result)

	// once can declare a var at the if stmt and it is available only during the scope if

	if y := 34; y > 0 {
		fmt.Println("The val is Greater than 0 for result2") 
	} else {
		fmt.Println("The val is Less than 0 for result2")
	}

}
