package main

import (
	"fmt"
	"time"
)

func main() {
	// the package infor at https://golang.org/pkg/time/
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("Time right now is: %s\n", &now)
	fmt.Println("The month is: ", t.Month())
	fmt.Println("The Day is: ", t.Day())
	fmt.Println("The Weekday is: ", t.Weekday())

	tmrw := t.AddDate(0, 0, 1)
	fmt.Printf("Tmrw is %v, %v, %v , %v\n", 
		tmrw.Weekday(), tmrw.Month(), tmrw.Day(), tmrw.Year())

	// in order to customize the date format we can define a new format and use it to display
	longFormat := "Monday, January 2, 2014"
	fmt.Println("Tmrw long Format is: ", tmrw.Format(longFormat))
	shFormat := "2016/28/04"
	fmt.Println("Tmrw is shFormat: ", shFormat)
}