package main

import (
	"fmt"
	"math"
)

func main() {
	primes := generatePrimes(10)

	fmt.Println("Prime numbers:")
	fmt.Println(primes)
}

func checkPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(n int) []int {
	primes := []int{}

	for i := 2; len(primes) < n; i++ {
		if checkPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
