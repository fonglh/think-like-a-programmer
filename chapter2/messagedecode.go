package chapter2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read the message and tokenize it.
func readMessage() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message as a string of comma separated numbers: ")
	inputString, _ := reader.ReadString('\n')
	//remove trailing newline
	inputString = inputString[:len(inputString)-1]

	return strings.Split(inputString, ",")
}

// Implement this for practise.
func stringToInt(s string) int {
	var number int
	for i := 0; i < len(s); i++ {
		number = number*10 + (int(s[i]) - '0')
	}
	return number
}

func convertMessageToIntSlice(inputSlice []string) []int {
	numbers := make([]int, len(inputSlice))

	for i, str := range inputSlice {
		numbers[i] = stringToInt(str)
	}
	return numbers
}

// 1 returns A, 2 returns B and so on...
func numToUpperChar(number int) string {
	char := number + 'A' - 1
	return string(char)
}

// 1 returns a, 2 returns b and so on...
func numToLowerChar(number int) string {
	char := number + 'a' - 1
	return string(char)
}

// return punctuation mark defined in problem
func numToPunctuation(number int) string {
	punctuationLookup := []string{" ", "!", "?", ",", ".", " ", ";", "\"", "'"}
	return punctuationLookup[number]
}

// increment mode modulo 3
// 0: uppercase
// 1: lowercase
// 2: punctuation
func nextMode(mode int) int {
	return (mode + 1) % 3
}

func DecodeMessage() {
	var output string
	var mode int

	// Get input from user
	inputSlice := readMessage()
	numberSlice := convertMessageToIntSlice(inputSlice)

	for _, number := range numberSlice {
		switch mode {
		case 0:
			number = number % 27
			if number == 0 {
				mode = nextMode(mode)
			} else {
				output += numToUpperChar(number)
			}
		case 1:
			number = number % 27
			if number == 0 {
				mode = nextMode(mode)
			} else {
				output += numToLowerChar(number)
			}
		case 2:
			number = number % 9
			if number == 0 {
				mode = nextMode(mode)
			} else {
				output += numToPunctuation(number)
			}
		}
	}

	fmt.Println(output)
}
