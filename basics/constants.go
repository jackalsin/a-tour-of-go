package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "驶接"
	fmt.Println("hello", world)
	fmt.Println("happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
