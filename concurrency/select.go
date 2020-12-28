package main

import "fmt"

func fab(val chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case val <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("result ", x)
			return
		}
	}
}
func main() {
	val := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-val)
		}
		quit <- 0
	}()
	fab(val, quit)
}
