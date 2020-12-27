package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m map[string]Vertex
	fmt.Println(m == nil) // true
	m = make(map[string]Vertex)
	fmt.Println(m)
	m["a"] = Vertex{1, 3}
	fmt.Println(m["a"])
}
