package main

import "fmt"

// Exercise 2: Create a program that uses goroutines and channels to calculate the sum of numbers from 1 to 100 concurrently.
func main() {
	ch := make(chan int)
	n := 100

	go task2(ch, n)
	fmt.Println(<-ch)

}

func task2(ch chan int, n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	ch <- sum
}
