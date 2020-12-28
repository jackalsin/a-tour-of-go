package main

import "fmt"

var id = 1

func sum(arr []int, c chan int) {
	curId := id
	id++
	fmt.Println(curId)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(curId)
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	left, right := <-c, <-c
	fmt.Println(left, right, left+right)
}
