package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{4, 2, 5, 1, 3}
	sortedArr := QuickSort(arr)
	fmt.Println(sortedArr) // Output: [1 2 3 4 5]
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Choose a random pivot index
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(arr))

	// Move the pivot element to the front of the array
	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]

	pivot := arr[0]
	var left []int
	var right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
			fmt.Println(right)
		}
	}

	leftSorted := QuickSort(left)
	rightSorted := QuickSort(right)

	return append(append(leftSorted, pivot), rightSorted...)
}
