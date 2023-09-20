package main

import (
	"fmt"
)

// Exercise 12: Create a program that uses goroutines and channels to find the sum of even numbers in a given slice concurrently.
func main() {
	ch := make(chan int)
	var n = []int{1, 2, 3, 4, 5, 6, 7, 8}

	go sumOfEven(ch, n)

	fmt.Println(<-ch)
}

func sumOfEven(ch chan int, nums []int) {
	sum := 0
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}
	ch <- sum
}
