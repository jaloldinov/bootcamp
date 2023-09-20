package main

import (
	"fmt"
)

func main() {
	var a, b, c float64 = 5, 4, 3

	prove(a, b, c)

}

func prove(a, b, c float64) {

	var num1, num2, num3, min, mid, max int

	fmt.Println("a = ")
	fmt.Scanln(&num1)
	fmt.Println("b = ")
	fmt.Scanln(&num2)
	fmt.Println("c = ")
	fmt.Scanln(&num3)

	if num1 >= num2 && num1 >= num3 {
		max = num1
		if num2 >= num3 {
			min = num3
		} else {
			min = num2
		}
	} else if num2 >= num1 && num2 >= num3 {
		max = num2
		if num1 >= num3 {
			min = num3
		} else {
			min = num1
		}
	} else {
		max = num3
		if num1 >= num2 {
			min = num2
		} else {
			min = num1
		}
	}

}
