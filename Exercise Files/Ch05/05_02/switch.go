package main

import (
	"fmt"
	"math/rand"
	"time"	
) 

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
		case 1:
			result = "Sunday"
		case 7:
			result = "Sunday"
		default:
			result = "Weekday"
	}

	fmt.Println("Day", dow, ",", result)

	// we can use the if condition within the switch stmt, also multiple stmts can be written with ; semi colon
	// fallthrough stmt can be used to execute the default stmt 
	x := -42;
	switch {
		case x < 0:
			result = "less than 0"
			// fallthrough
		case x == 0:
			result = "equal 0"
		default:
			result = "Greater than 0"
	}
	fmt.Println(result)
}
