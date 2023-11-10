package main

import "fmt"

func main() {

	// arr := []int{1, 2, 3, 1, 1, 3}
	arr := []int{1, 1, 1, 1}

	fmt.Println(findDup(arr))
}

func findDup(arr []int) int {

	var sum int

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				sum += 1
			}
		}
	}

	return sum

}
