package main

import (
	"fmt"
	"strings"
)

// Exercise 13: Create a program that uses goroutines and channels to find the length of the longest word in a given string concurrently.
func main() {
	text := "Hello everyone, my name is OMADBEKJALOLDINOV"
	ch := make(chan int)

	go findLongestWord(text, ch)

	longestLength := <-ch
	fmt.Printf("the length of the longest word in a given string: %d\n", longestLength)
}

func findLongestWord(text string, ch chan int) {
	words := strings.Fields(text)
	longestword := 0

	for _, word := range words {
		length := len(word)
		if length > longestword {
			longestword = length
		}
	}
	ch <- longestword
}
