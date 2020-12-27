package main

import "fmt"

func main() {
	var s []int
	fmt.Println(cap(s), len(s))
	if s == nil {
		fmt.Println("s is nil.")
	}
}
