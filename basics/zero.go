package main

import "fmt"

// this is to explain the default value

func main() {
	var i int
	var f float64
	var b bool
	var s string // default value is ""
	// specifier number -> v, string -> %q
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
