package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 3, 1, 2, 3, 3, 3, 546, 23, 3, 4, 54, 65, 76, 4}
	countNumber(arr)
}

func countNumber(arr []int) {

	// var counter int
	var container = make(map[int]int)

	for _, num := range arr {
		container[num]++
	}
	for num, count := range container {
		fmt.Printf("%d repeated: %d times\n", num, count)
	}

}
