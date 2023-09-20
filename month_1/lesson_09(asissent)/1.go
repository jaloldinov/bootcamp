package main

import "fmt"

func main() {
	command := "G()()()()(al)G()"
	// Output: "Gooooal"

	fmt.Println(interpret(command))
}

func interpret(command string) string {

	var result string

	for i := range command {
		if command[i] == 'G' {
			result += "G"
		} else if command[i] == '(' && command[i+1] == ')' {
			result += "o"
		} else if command[i] == '(' && command[i+1] == 'a' {
			result += "al"

		}
	}
	return result
}
