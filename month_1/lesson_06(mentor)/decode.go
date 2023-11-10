package main

import "fmt"

func main() {

	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"

	fmt.Println(decodeMessage(key, message)) // Output: "this is a secret"
}

func decodeMessage(key string, message string) string {
	mapString := make(map[rune]rune)
	idx := 0
	for _, v := range key {
		if v == ' ' {
			continue
		}

		if _, ok := mapString[v]; !ok {
			mapString[v] = rune(idx) + 'a'
			idx++
		}
	}

	res := []rune{}
	for _, v := range message {
		if v == ' ' {
			res = append(res, ' ')
		} else {
			res = append(res, mapString[v])
		}
	}

	return string(res)
}
