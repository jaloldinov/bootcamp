package staircase

import "fmt"

func Staircase(n int32) {

	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= n; j++ {
			if j <= n-i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
