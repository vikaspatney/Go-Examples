package main

import "fmt"

// Go supports interfaces to create high level abstraction, if a type implements all methods of the interface
// then it is the implementation of that interface; there is no keyword

type Animal interface{
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak () string {
	return "woof!"
}

type Cat struct {
}

func (c Cat) Speak () string {
	return "meow!"
}

type Cow struct {
}

func (c Cow) Speak () string {
	return "mooo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
