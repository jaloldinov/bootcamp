package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var playerName string
	fmt.Print("What's your name? ")
	fmt.Scan(&playerName)

	choices := map[int]string{
		0: "Scissors",
		1: "Paper",
		2: "Stone",
	}

	playerScore := 0
	computerScore := 0

	for playerScore < 3 && computerScore < 3 {

		var playerChoice int
		fmt.Printf("Player: %d Computer: %d\n", playerScore, computerScore)
		fmt.Printf("Enter your choice (0 => Scissors, 1 => Paper, 2 => Stone): ")
		fmt.Scan(&playerChoice)

		computerChoice := rand.Intn(3)

		fmt.Printf("%s chose %s.\n", playerName, choices[playerChoice])
		fmt.Printf("Computer chose %s.\n", choices[computerChoice])

		switch {
		case playerChoice == computerChoice:
			fmt.Println("It's a tie!")
		case playerChoice == 0 && computerChoice == 1:
			fmt.Printf("%s wins!\n", playerName)
			playerScore++
		case playerChoice == 1 && computerChoice == 2:
			fmt.Printf("%s wins!\n", playerName)
			playerScore++
		case playerChoice == 2 && computerChoice == 0:
			fmt.Printf("%s wins!\n", playerName)
			playerScore++
		default:
			fmt.Println("Computer wins!")
			computerScore++
		}
	}

	fmt.Printf("%s's final score: %d\n", playerName, playerScore)
	fmt.Printf("Computer's final score: %d\n", computerScore)
	if playerScore > computerScore {
		fmt.Printf("%s wins the game!\n", playerName)
	} else {
		fmt.Println("Computer wins the game!")
	}
}
