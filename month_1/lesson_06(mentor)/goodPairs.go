package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 1, 1, 3}
	// arr := []int{1, 1, 1, 1}

	fmt.Println(numIdenticalPairs(arr))
}

func numIdenticalPairs(nums []int) int {
	var count = 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}
