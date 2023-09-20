package main

import (
	"fmt"
)

func main() {
	var a interface{}
	a = 2

	// if _, ok := a.(string); ok {
	// 	fmt.Println("String")
	// } else if _, ok := a.(int); ok {
	// }

	switch a.(type) {
	case int:
		fmt.Println("an integer")
	case string:
		fmt.Println("a string")
	default:
		fmt.Println("an unknown type")
	}
}
