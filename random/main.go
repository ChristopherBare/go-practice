package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	for running {
		guessesTaken := 0
		randomNumber := rand.Intn(100)
		fmt.Println("Generating number between 1 and 100")
		for guessesTaken < 7 {
			fmt.Println("Guess the number")
			// Use scanner.Scan() to read the input
			if scanner.Scan() {
				// Use scanner.Text() to get the input as a string
				input := scanner.Text()

				value, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				guessesTaken += 1
				if value < randomNumber {
					fmt.Println("too low, try again")
				} else if value > randomNumber {
					fmt.Println("too high, try again")
				} else {
					fmt.Printf("You guessed the number correctly! The number is %d\n", randomNumber)
					if scanner.Scan() {
						// Use scanner.Text() to get the input as a string
						fmt.Println("Play again? (y/n)")
						input := scanner.Text()

						if input == "y" || input == "Y" {
							guessesTaken = 0
						} else {
							running = false
						}
					} else {
						fmt.Println("Error reading input:", scanner.Err())
					}
				}
			} else {
				fmt.Println("Error reading input:", scanner.Err())
			}
		}
		if guessesTaken >= 7 {
			fmt.Println("You are a failure.")
			fmt.Println("Play again? (y/n)")
			if scanner.Scan() {
				// Use scanner.Text() to get the input as a string
				input := scanner.Text()

				if input == "y" || input == "Y" {
					guessesTaken = 0
				} else {
					running = false
				}
			}
		}
	}
}
