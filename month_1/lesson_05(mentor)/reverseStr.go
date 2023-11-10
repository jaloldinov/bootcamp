package main

import "fmt"

func main() {

	str := "vonidlolaJ kebdamO"
	reverseString(str)
}

func reverseString(str string) {
	var reversed string

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	fmt.Println(reversed)
}
