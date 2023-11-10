package main

import (
	"fmt"
	"sort"
)

// Exercise 9: Create a program that uses goroutines and channels to sort a given slice of integers concurrently.
func main() {
	ch := make(chan []int)
	var n = []int{2, 1, 4, 3, 6, 5, 0, 7}

	go sorting(ch, n)

	fmt.Println(<-ch)
}

func sorting(ch chan []int, n []int) {
	sort.Ints(n)
	ch <- n
}
