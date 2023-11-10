package main

import (
	"fmt"
)

func main() {
	var arr = []int{5, 0, 1, 2, 3, 4}
	//    [4, 5, 0, 1, 2, 3]
	fmt.Println(changer(arr))
}

func changer(arr []int) []int {
	var newArr []int
	for _, v := range arr {
		newArr = append(newArr, arr[v])
	}
	return newArr
}
