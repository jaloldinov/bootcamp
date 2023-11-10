package main

import "fmt"

func main() {
	var nums = []int{8, 1, 2, 2, 3}

	fmt.Println(smallerNumbersThanCurrent(nums)) //[4,0,1,1,3]
}

func smallerNumbersThanCurrent(nums []int) []int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		arr = append(arr, count)
	}
	return arr
}
