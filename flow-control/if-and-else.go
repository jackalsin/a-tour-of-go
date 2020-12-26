package main

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v < 17 {
		fmt.Printf("17 >= %g\n", lim)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though -> 这是指else
	return lim
}

func main() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(4, 2, 20),
		pow2(3, 3, 20),
	)
}
