package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 2, 3, 7

	Arif(a, b, c)
	Geomet(a, b, c)
}

func Arif(a, b, c float64) {

	result := (a + b + c) / 3
	fmt.Println("Arifmetik = ", result)
}

func Geomet(a, b, c float64) {

	result := math.Pow((a * b * c), 1./3)
	fmt.Println("Gemometrik = ", result)
}
