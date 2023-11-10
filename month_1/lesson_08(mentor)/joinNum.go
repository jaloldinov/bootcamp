package main

import (
	"fmt"
)

func main() {
	a := 123
	b := 46789
	joinNum(a, b)
}

func joinNum(a, b int) {
	var counter int
	n := b

	for n > 0 {
		n = n / 10
		counter++
	}

	for i := 0; i < counter; i++ {
		a *= 10
	}
	a += b

	fmt.Println(a)
}
