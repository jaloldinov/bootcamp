package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 6, 8, 9, 11}
	low := 6
	high := 15

	fmt.Println(lowAndHigh(arr, low, high))

}

func lowAndHigh(arr []int, low, high int) []int {

	result := make([]int, 0)
	for _, num := range arr {
		if num >= low && num <= high {
			result = append(result, num)
		}
	}
	return result
}
