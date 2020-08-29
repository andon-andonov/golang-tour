package main

import (
	"fmt"; 
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 0.0000000001
	for {
		pz := z
		z -= (z*z - x) / (2*z)
		if math.Abs(pz - z) <= eps {
			break
		}
	}
	return z
}

func main() {
	x := 1.61
	fmt.Println(math.Sqrt(x))
	fmt.Println(Sqrt(x))
}
