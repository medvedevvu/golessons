package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	i := 1
	z := 1.0
	for i <= 10 {
		z = z - (z*z-x)/(2*z)
		i += 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
