package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	 z  := 1.0
	for i:=1; i <= 100; i++ {
		z = z - ( (z*z) - x ) / (2*z) 
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
