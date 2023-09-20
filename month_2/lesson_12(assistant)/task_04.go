package main

import (
	"fmt"
	"math"
)

// Exercise 4: Create a program that uses goroutines and channels to check if a given number is prime concurrently.
func main() {
	ch := make(chan bool)
	n := 7

	go checkPrime(ch, n)
	if <-ch {
		fmt.Printf("%d is prime\n", n)
	} else {
		fmt.Printf("%d is not prime\n", n)
	}
}

func checkPrime(ch chan bool, n int) {
	if n <= 1 {
		ch <- false
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			ch <- false
			return
		}
	}
	ch <- true
}
