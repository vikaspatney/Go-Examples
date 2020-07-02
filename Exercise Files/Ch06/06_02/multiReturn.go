package main

import "fmt"

func main() {

	n1, l1 := FullName("vik", "pat")
	fmt.Printf("Fullname: %v, no of chars: %v\n\n", n1, l1)

	n1, l1 = FullnameRetParam("vik", "pat")
	fmt.Printf("Fullname RET: %v, no of chars: %v", n1, l1)
}

func FullName (f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// we can also declare the returning variable in order at the func decleration 
func FullnameRetParam(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}