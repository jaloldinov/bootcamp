package main

import "fmt"

func main() {
	words := []string{"cd", "ac", "dc", "ca", "zz"}
	// Output: "Gooooal"

	fmt.Println(maximumNumberOfStringPairs(words))
}

func maximumNumberOfStringPairs(words []string) int {
	pairsCounter := make(map[string]struct{})
	var counter int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word[0] > word[1] {
			word = string([]byte{word[1], word[0]})
		}

		if _, ok := pairsCounter[word]; ok {
			counter++
		}

		pairsCounter[word] = struct{}{}
	}

	return counter
}
