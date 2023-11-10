package main

import "fmt"

func main() {

	arr := [][]int{{1, 2, 3}, {2, 5, 6}, {7, 8, 9}, {3, 1, 2}}
	fmt.Println(compareSlices(arr))

}

func compareSlices(arr [][]int) int {
	maxSum := 0

	for _, subSlice := range arr {
		sum := 0
		for _, num := range subSlice {
			sum += num
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
