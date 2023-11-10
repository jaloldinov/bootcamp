package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	n := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(n) + 1
	fmt.Println("secret number: ", secretNumber)
	numGuesses := 1
	maxGuesses := int(math.Log2(float64(n))) + 2

	fmt.Printf("1 dan %d gacha son kiriting. Sizda %d ta urinish bor. Topolasizmi?\n", n, maxGuesses)
	for numGuesses < maxGuesses {
		var guess int
		fmt.Printf("Sizda %d ta urunish qoldi. tahminiz: ", maxGuesses-numGuesses)
		fmt.Scanln(&guess)

		if guess < secretNumber {
			fmt.Println("Juda kichik son kiritdingiz!")
		} else if guess > secretNumber {
			fmt.Println("Juda katta son kiritdingiz!")
		} else {
			fmt.Printf("Shapaloqlar bo'lsin! Siz %d marta urinishda toptingiz.\n", numGuesses)
			return
		}
	}

	fmt.Println("ooh no, topolmadingiz.  sirli son: ", secretNumber)
}
