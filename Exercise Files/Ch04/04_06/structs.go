package main

import (
	"fmt"
)

type Dog struct{
	Breed string // keep the first letter to be capital for making it a Public field
	Weight int 
}

func main() {

	poodle := Dog{"Poo", 34} // we use curly braces to initialize the object as there is no constructor associated
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	// we can also use the . operator in order to access the fields in the struct
	fmt.Printf("Breed: %v\nWeight: %v", poodle.Breed, poodle.Weight)
}