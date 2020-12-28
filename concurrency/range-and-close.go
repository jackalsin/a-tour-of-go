package main

import "fmt"

func fab(len int, c chan int) {
	next1, next2 := 0, 1
	for i := 0; i < len; i++ {
		c <- next1
		next1, next2 = next2, next1+next2
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fab(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}
