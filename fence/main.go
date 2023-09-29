package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	running := true
	for running {
		var fenceLen float32
		var fenceWidth float32
		var totalGallons int
		var totalSqft int
		pricePerGallon := 29.85
		sqftPerGallon := 422

		scanner := bufio.NewScanner(os.Stdin)

		// Use scanner.Scan() to read the input
		if scanner.Scan() {
			// Use scanner.Text() to get the input as a string
			fmt.Println("fence length")
			input := scanner.Text()
			fmt.Println("fence width")
			input2 := scanner.Text()

			// Convert the input string to a float32
			value, err := strconv.ParseFloat(input, 32)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			value2, err := strconv.ParseFloat(input2, 32)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fenceLen = float32(value)
			fenceWidth = float32(value2)

			totalSqft = int(fenceLen * fenceWidth)
			totalGallons = totalSqft / sqftPerGallon

			if totalGallons < 1 {
				totalGallons = 1
			}
			if totalGallons%2 != 0 && totalGallons > 1 {
				totalGallons += 1
			}
			totalPrice := float64(totalGallons) * pricePerGallon
			fmt.Sprintf("total gallons %d", totalGallons)
			fmt.Sprintf("total sqft %d", totalSqft)
			fmt.Sprintf("total price %d", totalPrice)
			fmt.Println("play again?(y/n)")
			input3 := scanner.Text()
			if input3 == "n" || input3 == "N" {
				running = false
			}
		} else {
			fmt.Println("Error reading input:", scanner.Err())
		}
	}
}
