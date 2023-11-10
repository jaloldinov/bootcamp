package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "bir ikki bir uch besh bir ikki on qalesan salom salom "
	countNumber(str)
}

func countNumber(str string) {

	splited := strings.Fields(str)
	var container = make(map[string]int)

	for _, word := range splited {
		container[word]++
	}
	for word, count := range container {
		fmt.Printf("%s repeated: %d times\n", word, count)
	}
}
