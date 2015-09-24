package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func printCurrent(digits int) string {
	var s string
	for i := 1; i <= digits; i++ {
		s = s + "*"
	}
	return s
}

func makeGuess() int {
	fmt.Print("Make a guess:   ")
	var guess int
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return guess
}

func analyzeGuess(mystery int, guess int, choices []string) (string, string, []string) {
	var sCurr string
	var sHint string
	sMystery := strconv.Itoa(mystery)
	sGuess := strconv.Itoa(guess)

	for i, c := range sGuess {
		num, _ := strconv.Atoi(string(c))
		if !strings.Contains(sMystery, string(c)) {
			sCurr = sCurr + "B"
			choices[num] = "*"
			sHint = sHint + "*"
		} else {
			if string(sMystery[i]) == string(c) {
				sCurr = sCurr + "F"
				sHint = sHint + string(c)

			} else {
				sCurr = sCurr + "P"
				sHint = sHint + "*"
			}
		}
	}
	return sCurr, sHint, choices
}

func checkForWin(current string) bool {
	for _, c := range current {
		if string(c) != "F" {
			return false
		}
	}
	return true
}

func printInstructions() {
	fmt.Println()
	fmt.Println("Welcome! Let's play...")
	fmt.Println()
	fmt.Println(" ____  _           _____                   _ ____                   _           ")
	fmt.Println("|  _ \\(_) ___ ___ |  ___|__ _ __ _ __ ___ (_) __ )  __ _  __ _  ___| |___      ")
	fmt.Println("| |_) | |/ __/ _ \\| |_ / _ \\ '__| '_ ` _ \\| |  _ \\ / _` |/ _` |/ _ \\ / __| ")
	fmt.Println("|  __/| | (_| (_) |  _|  __/ |  | | | | | | | |_) | (_| | (_| |  __/ \\__ \\    ")
	fmt.Println("|_|   |_|\\___\\___/|_|  \\___|_|  |_| |_| |_|_|____/ \\__,_|\\__, |\\___|_|___/")
	fmt.Println("                                                         |___/                  ")
	fmt.Println()
	fmt.Println("Here are the rules: ")
	fmt.Println("You will choose the number of digits to guess in the mystery number.")
	fmt.Println("The computer will randomly choose the mystery number.")
	fmt.Println("You will attempt to guess the mystery number.")
	fmt.Println("For each guess, you will receive a hint from the following:")
	fmt.Println()
	fmt.Println("	P = Pico  - the digit is correct but not in this place value")
	fmt.Println("	F = Fermi - the digit is correct and in the correct place value")
	fmt.Println("	B = Bagel - the digit is incorrect and doesn't appear anywhere in the number")
	fmt.Println()
	fmt.Println("Good luck! And at anytime you may type 'Ctrl+C' to quit.")
	fmt.Println()
	fmt.Println()
}

func pickNumDigits() int {
	var digits int
	valid := false
	for !valid {
		fmt.Print("Please pick the number of digits in the mystery number: ")
		_, err := fmt.Scanf("%d", &digits)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if digits <= 0 {
			fmt.Println("Number of digits must be greater than zero")
			continue
		}
		valid = true
	}
	return digits
}

func main() {

	printInstructions()

	// infinite loop unless player decides to quit
	for {
		choices := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		digits := pickNumDigits()

		rand.Seed(time.Now().UTC().UnixNano())
		min := int(math.Pow10(digits - 1))
		max := int(math.Pow10(digits)) - 1
		mystery := rand.Intn(max)

		s := printCurrent(digits)
		fmt.Printf("Mystery number: %s   ==>   Choices: %s\n", s, choices)

		numGuesses := 0
		correct := false
		for correct != true {
			numGuesses = numGuesses + 1
			guess := makeGuess()
			if guess < min || guess > max {
				fmt.Printf("Guess must be between %d and %d. Try again.\n", min, max)
				continue
			}

			current, hint, choices := analyzeGuess(mystery, guess, choices)
			fmt.Println("Hint:           " + current)
			fmt.Printf("Mystery number: %s  ==>   Choices: %s\n", hint, choices)

			if checkForWin(current) {
				fmt.Printf("\n *** CONGRATULATIONS!!! YOU GOT IT IN %d GUESSES ***\n\n", numGuesses)
				fmt.Println()
				fmt.Println("\nPress 'q' to quit or any other key to play again.")
				var response string
				_, err := fmt.Scanf("%s", &response)
				if (err != nil && err.Error() != "unexpected newline") || strings.HasPrefix(strings.ToLower(response), "q") {
					os.Exit(0)
				} else {
					break
				}
			}
		}
	}
}
