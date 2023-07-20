package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// oddEven()
	// swapTwoNumbers()
	// sumOfMinAndMax()
	// sumOfMinAndMax2()
	// distanceTwoPoints()
	solutionQuadraticEquation()
}

// odd or even
func oddEven() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println(num, "is even number")
	} else {
		fmt.Println(num, "is odd number")
	}
}

// swap two numbers
func swapTwoNumbers() {
	var a int
	var b int
	fmt.Print("Enter a value for a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter a value for b: ")
	fmt.Scanln(&b)

	fmt.Println("Before swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a

	fmt.Println("After swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

// summ of maximum and minimum of given 3 numbers
func sumOfMinAndMax() {
	// var nums [3]int
	// fmt.Println("Enter a number for num 1: ")
	// fmt.Scanln(&nums[0])
	// fmt.Println("Enter a number for num 2: ")
	// fmt.Scanln(&nums[1])
	// fmt.Println("Enter a number for num 3: ")
	// fmt.Scanln(&nums[2])
	nums := [3]int{7, 1, 4}
	sort.Ints(nums[:])

	min := nums[0]
	max := nums[2]

	sum := min + max
	fmt.Println("sum of min and max of three numbers is =>", sum)
}

// summ of maximum and minimum of given 3 numbers
func sumOfMinAndMax2() {
	var num1, num2, num3, min, max int

	fmt.Println("Enter a number for num 1: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter a number for num 2: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter a number for num 3: ")
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

	sum := min + max

	fmt.Println("sum of min and max of three numbers is =>", sum)
}

func distanceTwoPoints() {
	var x1, y1, x2, y2 float64

	fmt.Print("Enter x1 point: ")
	fmt.Scanln(&x1)
	fmt.Print("Enter y1 point: ")
	fmt.Scanln(&y1)
	fmt.Print("Enter x2 point: ")
	fmt.Scanln(&x2)
	fmt.Print("Enter y2 point: ")
	fmt.Scanln(&y2)

	// d=√((x2-x1)²+(y2-y1)²)
	// distance^2 = (x2 - x1)^2 + (y2 - y1)^2
	// Calculate x^y – Pow() in Go (Golang)
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	fmt.Println(distance)
}

func solutionQuadraticEquation() {
	var a, b, c float64
	a, b, c = 1, 2, 3

	// discriminant
	discriminant := (math.Pow(b, 2) - (4 * a * c))

	// x1 and x2  (-b -+ dis ) / 2*a
	x1 := (-b - discriminant) / (2 * a)
	x2 := (-b + discriminant) / (2 * a)

	fmt.Println("x1 and x2: ", x1, x2)
}
