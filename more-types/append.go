package main

import "fmt"

func main() {
	//original()
	sliceReference()
}

func original() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	fmt.Println("\nNew slice test.")
	thisIsArray := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(thisIsArray)
	var thisIsSlice []int = []int{0, 1, 2, 3, 4, 5}
	var slice []int = thisIsSlice[1:5]
	printSlice(slice)
	slice1 := append(slice, 6)
	printSlice(slice1)
	printSlice(thisIsSlice)

	fmt.Println("\nappend 7 to slice2")
	slice2 := append(slice1, 7)
	printSlice(slice1)
	printSlice(slice2)
	printSlice(thisIsSlice)

	fmt.Println("\nupdate thisIsSlice[2] = 23")
	thisIsSlice2 := thisIsSlice
	thisIsSlice[2] = 23
	printSlice(slice1)
	printSlice(slice2)
	printSlice(thisIsSlice)
	printSlice(thisIsSlice2)
}

func sliceReference() {
	slice1 := []int{0, 1, 2}
	printSlice(slice1)
	fmt.Println("Appending 3 to slice1")
	slice2 := append(slice1, 3)
	printSlice(slice2) // 此时len + 1， cap 翻倍
	fmt.Println("slice1[0] = -1")
	slice1[0] = -1
	printSlice(slice1) // [-1 1 2]
	printSlice(slice2) // [0 1 2 3]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
