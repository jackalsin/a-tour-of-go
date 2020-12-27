package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	val, ok := m["a"]
	fmt.Println(val, ok)
	delete(m, "b")
	fmt.Println(m["a"])
	delete(m, "a")
	fmt.Println(m["a"])
}
