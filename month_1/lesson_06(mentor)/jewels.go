package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"

	fmt.Println(numJewelsInStones(jewels, stones))
}

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				count++
			}
		}
	}
	return count
}
