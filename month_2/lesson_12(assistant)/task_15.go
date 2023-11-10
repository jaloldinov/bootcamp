package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Exercise 15: Create a program that uses goroutines
// and channels to implement a simple producer-consumer pattern. The producer goroutine should generate numbers
// and send them to the consumer goroutine through a channel, which will print the received numbers.
func main() {
	ch := make(chan int)
	fmt.Println("Program started")

	go producer(ch)
	consumer(ch)

	fmt.Println("Program finished")
}

func producer(ch chan<- int) {
	// random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(200 * time.Millisecond)
	}

	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
