package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkHelper(t.Left, ch)
	ch <- t.Value
	walkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else if ok1 {
			return false
		} else if ok2 {
			return false
		} else {
			return true
		}
	}
}

func main() {
	tree1, tree2 := tree.New(1), tree.New(2)
	fmt.Println(Same(tree1, tree2))
}
