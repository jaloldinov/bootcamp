package main

import (
	"fmt"
)

// Exercise 6: Create a program that uses goroutines and channels to reverse a string concurrently.
func main() {
	ch := make(chan string)
	str := "Omadbek Jaloldinov"

	go reverseString(ch, str)

	fmt.Println(<-ch)
}

func reverseString(ch chan string, str string) {
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	ch <- reversed
}
