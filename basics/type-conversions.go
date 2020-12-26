package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	// shrink to zero
	var twoFive, threeFive float64 = float64(2.5), float64(3.5)
	fmt.Println(
		int(twoFive),
		int(threeFive),
		int(-twoFive),
		int(-threeFive),
	)
}
