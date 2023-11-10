package main

import "fmt"

// Exercise 3: Create a program that uses goroutines and channels to find the factorial of a given number concurrently.
func main() {
	ch := make(chan int)
	n := 10

	go task3(ch, n)
	fmt.Println(<-ch)

}

func task3(ch chan int, n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	ch <- sum
}
