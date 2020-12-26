package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return add(x, y)
}

// no overriding
//func add(x int, y int, b float64) int {
//	return int(b) + add(x, y)
//}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add2(3, 4))
}
