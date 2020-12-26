package main

import "fmt"

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2} // has type *Vertex
	p2 = &p
)

func main() {
	fmt.Println(v1, p, v2, v3, p2)
}
