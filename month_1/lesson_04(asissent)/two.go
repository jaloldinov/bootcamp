package main

import "fmt"

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	fmt.Println(interleaveHalves(arr, 3))
}

func interleaveHalves(arr []int, n int) []int {

	firstHalf := arr[:n]
	secondHalf := arr[n:]
	result := []int{}

	for i := 0; i < n; i++ {

		result = append(result, firstHalf[i], secondHalf[i])
	}

	return result
}
