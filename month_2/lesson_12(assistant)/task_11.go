package main

import "fmt"

// Exercise 11: Create a program that uses goroutines and channels to calculate the Fibonacci sequence up to a given number concurrently.

func main() {

	ch := make(chan int)
	n := 10

	go fibonacci(ch, n)

	for v := range ch {
		if v <= n {
			fmt.Println(v)
		}
	}

}

func fibonacci(ch chan int, n int) {
	if n <= 0 {
		close(ch)
		return
	}
	first, second := 0, 1
	for i := 0; i < n; i++ {
		ch <- first

		first, second = second, first+second
	}
	close(ch)
}
