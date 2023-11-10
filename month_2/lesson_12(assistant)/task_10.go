package main

import "fmt"

// Exercise 10: Create a program that uses goroutines and channels to find the sum of squares of numbers from 1 to 10 concurrently.
func main() {
	ch := make(chan int)
	n := 10

	go findSquare(ch, n)

	fmt.Println(<-ch)
}

func findSquare(ch chan int, n int) {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	ch <- sum
}
