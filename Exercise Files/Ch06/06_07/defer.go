package main

import (
	"fmt"
)

func main() {
	// the defered stmt uses the LIFO for treating the defered stmt. the last defer stmt is first to be displayed below
	// defer stmt is treated at the end of the function, not the application
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	defer fmt.Println("Stmt 1")
	defer fmt.Println("Stmt 2")
	myFunc()
	defer fmt.Println("Stmt 3")
	defer fmt.Println("Stmt 4")
	fmt.Println("undefered stmt")

	x := 1000
	defer fmt.Println("Value of x: ", x)
	x++
	fmt.Println("Value of x without defer: ", x)

}

func myFunc() {
	defer fmt.Println("defered in the function")
	fmt.Println("not defered in the function")
}
