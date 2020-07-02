package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {
	
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)
	
	f1, f2, f3 := 23.4, 21.4, 67.5
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is: ", floatSum)

	var b1, b2, b3, bsum big.Float 
	b1.SetFloat64(23.5)
	b2.SetFloat64(43.2)
	b3.SetFloat64(92.9)

	bsum.Add(&b1, &b2).Add(&bsum, &b3)
	// if we dont add the & sign then it will display the whole complex number instead
	fmt.Printf("displaying the bsum val:=%.10g\n ", &bsum)

	// math package https://golang.org/pkg/math/
	circleRad := 15.5
	circum := circleRad * math.Pi
	fmt.Printf("circumfrence: %.2f\n", circum)
}
