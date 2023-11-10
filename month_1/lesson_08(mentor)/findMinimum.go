package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 1, 8, -2, 3, 4, -2, 8}
	countNumber(arr)
}

func countNumber(arr []int) {

	var counter int = 1
	var min int = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			counter = 1
		} else if arr[i] == min {
			counter++
		}
	}

	fmt.Println("minimum:", min)
	fmt.Println("repeated:", counter)
}
