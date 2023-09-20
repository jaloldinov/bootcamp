package main

import (
	"BOOTCAMP/backend_bootcamp_19_07_2023/candles"
	comparetriplets "BOOTCAMP/backend_bootcamp_19_07_2023/compareTriplets"
	"BOOTCAMP/backend_bootcamp_19_07_2023/fibonacci"
	reversenumber "BOOTCAMP/backend_bootcamp_19_07_2023/reverseNumber"
	"BOOTCAMP/backend_bootcamp_19_07_2023/staircase"
	"fmt"
)

func main() {

	// ### 1. Compare the Triplets
	fmt.Println(compareTripletsPrint())

	// ### 2. Staircase
	staircase.Staircase(6)

	// ### 3. Birthday Cake Candles
	fmt.Println(candlesPrint())

	// Reverse given number (67243 -> 34276)
	fmt.Println(reversenumber.ReverseNumber(67243))

	// ### 4. Fibonacci Number
	fmt.Println(fibonacci.Fib(10))
}

func compareTripletsPrint() []int32 {
	var alice = []int32{5, 6, 7}
	var bob = []int32{3, 6, 10}
	return comparetriplets.CompareTriplets(alice, bob)
}

func candlesPrint() int32 {
	var arr = []int32{4, 4, 3, 1, 3, 5, 67, 8, 67}
	return candles.BirthdayCakeCandles(arr)
}
