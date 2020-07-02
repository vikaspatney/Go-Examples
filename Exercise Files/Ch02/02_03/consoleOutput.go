package main

import "fmt"

func main() {
	
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true
	strlen, err := fmt.Println(str1, str2, str3)

	if err == nil{
		fmt.Println("String len:", strlen)
	}

	// using the Printf
	fmt.Printf("Value of the aNumber: %v\n", aNumber)
	fmt.Printf("Value of the isTrue: %v\n", isTrue)
	fmt.Printf("Value of the aNumber as a float64: %.2f\n", float64(aNumber))
	fmt.Printf("Data types for the vars: %T, %T, %T, %T, and %T\n",
		str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("Data types for the vars String: %T, %T, %T, %T, and %T",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
