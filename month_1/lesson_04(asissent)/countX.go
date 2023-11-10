package main

import "fmt"

func main() {

	arr := []string{"++x", "--x", "++x", "--x", "--x", "d"}

	fmt.Println(add(arr))
}

func add(arr []string) int {

	var sum int
	for _, v := range arr {
		if v == "++x" {
			sum++
		} else if v == "--x" {
			sum--
		} else {
			fmt.Println("wrong input !")
		}
	}

	return sum
}
