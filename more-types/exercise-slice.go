package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("dx = ", dx, ", dy = ", dy)
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i * j)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
