package main

import (
	"fmt"
)

// Exercise 8: Create a program that uses goroutines and channels to find the average of numbers in a given slice concurrently.\
func main() {
	ch := make(chan int)
	var n = []int{1, 2, 3, 4, 5}

	go average(ch, n)

	fmt.Println(<-ch)
}

func average(ch chan int, n []int) {
	sum := 0
	length := len(n)

	for _, v := range n {
		sum += v
	}

	ch <- int(sum / length)
}
