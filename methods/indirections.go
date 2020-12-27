package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// 虽然receiver是pointer，call用的value，但是会自动call pointer
	v.Scale(2)
	ScaleFunc(&v, 10) // {60, 80}

	p := &Vertex{4, 3}
	p.Scale(3) // {12, 9}
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
