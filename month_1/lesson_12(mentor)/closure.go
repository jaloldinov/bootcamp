package main

import "fmt"

func createAdder() func(int) int {
	sum := 0

	adder := func(num int) int {
		sum += num
		return sum
	}
	return adder
}

func main() {
	adder := createAdder()

	fmt.Println(adder(3))
	fmt.Println(adder(7))
	fmt.Println(adder(10))

}
