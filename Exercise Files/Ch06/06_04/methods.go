package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

// for creating member functions we need to define as below passing an instance or obj
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// passing the object as a pointer in the parameters
func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {

	poodle := Dog {"Poodle", 34, "Woof"}
	fmt.Println(poodle)
	// the member func can be accessed using the . operator
	poodle.Speak()

	// we can assign a value to the object directly using the . operator
	poodle.Sound = "Arf!"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()

}
