package main

import "fmt"

func main() {

	var num float64
	fmt.Print("num = ")
	fmt.Scanln(&num)

	chorak(num)
}

func chorak(num float64) {

	if num > 0 && num < 90 {
		fmt.Println("birinchi chorak")
	} else if num > 90 && num < 180 {
		fmt.Println("ikkinchi chorak")
	} else if num > 180 && num < 1270 {
		fmt.Println("uchinchi chorak")
	} else {
		fmt.Println("to'rtinchi chorak")
	}

}
