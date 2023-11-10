package main

import (
	"fmt"
)

// Exercise 5: Create a program that uses goroutines and channels to find the largest number in a given slice concurrently.
func main() {
	ch := make(chan int)
	var n = []int{1, 2, 41, 3, 4, 5, 6, 6, 7, 40, 42}

	go findLargest(ch, n)

	fmt.Printf("The largest number in the slice is %d\n", <-ch)
}

func findLargest(ch chan int, n []int) {
	largest := n[0]

	for _, num := range n {
		if num > largest {
			largest = num
		}
	}
	ch <- largest
}
