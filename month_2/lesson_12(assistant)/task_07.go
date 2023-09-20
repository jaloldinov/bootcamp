package main

import (
	"fmt"
)

// Exercise 7: Create a program that uses goroutines and channels to multiply each element in a given slice by 2 concurrently.
func main() {
	ch := make(chan []int)
	var n = []int{1, 2, 3, 4, 5}

	go dublicate(ch, n)

	fmt.Println(<-ch)
}

func dublicate(ch chan []int, n []int) {
	for i := range n {
		n[i] *= 2
	}

	ch <- n
}
