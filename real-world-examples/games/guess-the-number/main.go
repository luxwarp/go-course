package main

import (
	"fmt"
	"math/rand"
	"time"
)

const minNumber int = 1

var maxNumber int
var randomNumber int

func main() {
	fmt.Println("Guess the number")
	fmt.Printf("Enter a number above 0 that you want to guess for: ")
	fmt.Scan(&maxNumber)
	fmt.Println("Okey, now it's time to guess what number between 1 and", maxNumber, "that I have chosen.")
	rand.Seed(time.Now().UnixNano())

	randomNumber = rand.Intn(maxNumber-minNumber+1) + minNumber

	var guessedNumber int
	for guessedNumber != randomNumber {
		fmt.Printf("Guess: ")
		fmt.Scanln(&guessedNumber)

		if guessedNumber < randomNumber {
			fmt.Println("To low. Try again!")
		} else if guessedNumber > randomNumber {
			fmt.Println("To high. Try again!")
		} else {
			fmt.Println("Nice job! That's correct. You win!")
			return
		}
	}
}
