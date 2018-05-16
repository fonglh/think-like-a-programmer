package chapter2

// Exercise 2-6

import (
	"fmt"
	"strconv"
)

func numSliceToDecimal(digits []int) int {
	decimal := 0

	for _, digit := range digits {
		decimal = decimal*10 + digit
	}
	return decimal
}

func BinaryConvert() {
	numbers := readInput("Enter number to convert to binary: ")
	decimalInput := numSliceToDecimal(numbers)
	binaryString := ""

	for decimalInput > 0 {
		remainder := decimalInput % 2
		binaryString = strconv.Itoa(remainder) + binaryString
		decimalInput = decimalInput / 2
	}

	fmt.Println(binaryString)
}
