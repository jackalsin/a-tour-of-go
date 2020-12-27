package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior
//as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value
//declared in the main function.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
