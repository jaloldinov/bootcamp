package main

import "fmt"

// Exercise 1: Create a program that uses goroutines and channels to print numbers from 1 to 10 concurrently.
func main() {
	ch := make(chan int)
	n := 10

	go numberSender(ch, 0, n/2)
	go numberSender(ch, n/2, n)

	for {
		fmt.Println(<-ch)
	}
}

func numberSender(ch chan int, x, n int) {

	for i := x; i <= n; i++ {
		ch <- i
	}
}
