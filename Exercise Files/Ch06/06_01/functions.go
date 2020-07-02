package main

import "fmt"

func main() {
	doSomething()
	sum := addVal(23, 44)
	fmt.Println("the sum is: ", sum)

	sum = addAllVal(23, 45,54)
	fmt.Println("New sum :", sum)

}
// function overloading is not allowed, every func must have correct number of arguments
// lower case character makes the func private
func doSomething(){
	fmt.Println("Doing something")
}

// if both or all parameters have same data type then we can declare the type only once at the end
func addVal(val1 int, val2 int) int {
	return val1 + val2
}

func addAllVal(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}