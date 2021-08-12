package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)

	ch <- t.Value

	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}

	return true
}

func runEquivalentBinaryTreesExercise() {
	if Same(tree.New(1), tree.New(1)) != true {
		fmt.Println("FAILED: Same(tree.New(1), tree.New(1)) should be true not false")
	} else if Same(tree.New(1), tree.New(2)) != false {
		fmt.Println("FAILED: Same(tree.New(1), tree.New(2)) should be false not true")
	} else {
		fmt.Println("PASSED")
	}
}
