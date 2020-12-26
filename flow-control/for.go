package main

import "fmt"

func main() {
	sum := 0
	// ++i is illegal
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
