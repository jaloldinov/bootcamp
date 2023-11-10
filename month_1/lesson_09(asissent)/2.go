package main

import (
	"fmt"
)

func main() {
	str := "Hello"

	fmt.Println(toLowerCase(str))
}

func toLowerCase(s string) string {
	res := []byte(s)
	for i := 0; i < len(res); i++ {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
		fmt.Println(32)
	}
	return string(res)
}

// easy wayyyy
// func toLowerCase(s string) string {
// 	return strings.ToLower(s)
// }
