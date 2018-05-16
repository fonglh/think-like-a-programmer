package chapter2

import (
	"bufio"
	"fmt"
	"os"
)

func sumDoubledDigit(digit int) int {
	doubledDigit := digit * 2
	if doubledDigit >= 10 {
		return 1 + doubledDigit%10
	} else {
		return doubledDigit
	}
}

// Read input from the user, then return a slice of the individual digits as integers.
func readInput(prompt string) []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	inputString, _ := reader.ReadString('\n')

	numbers := make([]int, len(inputString)-1)
	for i := 0; i < len(inputString)-1; i++ {
		numbers[i] = int(inputString[i]) - 48
	}
	return numbers
}

func calculateChecksum(numbers []int) int {
	// Pretend we don't know how long `numbers` is at the start.
	var oddChecksum int
	var evenChecksum int

	// Calculate checksums for both odd and even length numbers.
	// Odd length doubles the 1st digit from the left, then every other digit.
	// Even length doubles the 2nd digit from the left, then every other digit.
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			oddChecksum += numbers[i]
			evenChecksum += sumDoubledDigit(numbers[i])
		} else {
			oddChecksum += sumDoubledDigit(numbers[i])
			evenChecksum += numbers[i]
		}
	}

	if len(numbers)%2 == 0 {
		return evenChecksum
	} else {
		return oddChecksum
	}
}

func ValidateLuhnChecksum() {
	numbers := readInput("Enter number for Luhn Checksum Validation: ")
	checksum := calculateChecksum(numbers)

	if checksum%10 == 0 {
		fmt.Println("Checksum is divisible by 10. Number is valid")
	} else {
		fmt.Println("Checksum is not divisible by 10. Number is invalid")
	}
}
