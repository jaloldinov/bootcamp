package main

import "fmt"

func main() {
	// factorial
	n := 10
	var result = make(chan int)
	// go factorial(n, result)
	// fact := <-result
	// fmt.Printf("Factorial of %d is: %d\n", n, fact)

	// fibonaci
	go fibonacci(n, result)
	for v := range result {
		if v <= n {
			fmt.Println(v)
		}
	}
}

func factorial(n int, result chan int) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	result <- fact
}

func fibonacci(n int, result chan int) {
	if n <= 0 {
		close(result)
		return
	}
	first, second := 0, 1
	for i := 0; i < n; i++ {
		result <- first

		first, second = second, first+second
	}
	close(result)
}
