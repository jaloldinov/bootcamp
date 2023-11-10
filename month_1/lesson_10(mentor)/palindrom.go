package main

import "fmt"

func main() {
	arr := []int{1, 2, 7, 2, 1} // true
	// arr := []int{1, 4, 2, 1} // false

	fmt.Println(chekPalindrom(arr))

}

func chekPalindrom(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			// fmt.Println(arr[i], arr[len(arr)-1-i])
			return false
		}
	}
	return true
}
