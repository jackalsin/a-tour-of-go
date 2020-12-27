package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := -1
	prev1, prev2 := 0, 1
	return func() int {
		i++
		if i == 0 {
			return 0
		} else if i == 1 {
			return 1
		} else {
			res := prev1 + prev2
			prev1 = prev2
			prev2 = res
			return res
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
