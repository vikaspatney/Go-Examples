package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly decl string"
	fmt.Println("Here is my implicit string: ", str1)
	// to know the value and type using Printf
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "An explicitly decl string"
	fmt.Println("Here is my explicit string: ", str2)
	// to know the value and type using Printf
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	// Title is used to convert first letter of every word as Capital
	fmt.Println(strings.Title(str1))

	// comparing the string case sensitive
	lValue := "hello"
	uVal := "HELLO"
	fmt.Println("are equal strings?: ", (lValue == uVal))
	// EqualFold(param1, param2) allows us to compare values as non case sensitive  
	fmt.Println("equal non case sensitive compare: ", strings.EqualFold(lValue, uVal))
	// contains (str,exp) allows us to secarch for a experssion within a given string
	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))

	// https://golang.org/pkg/strings/#example_Builder for complete ref
	// using the builder
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d\n...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
