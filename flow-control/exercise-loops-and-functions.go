package main

import (
	"fmt"
	"math"
)

const bias = 1e-3

func Sqrt(x float64) float64 {
	z := 1.0
	count := 1
	for {
		count += 1
		prevZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(prevZ-z) < bias {
			break
		}
	}
	fmt.Printf("Running count = %v\n", count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
