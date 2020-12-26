package main

import "fmt"

func main() {
	//				 0, 1, 2, 3,  4,  5
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
