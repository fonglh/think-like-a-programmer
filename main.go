package main

import (
	"fmt"

	"github.com/fonglh/think-like-a-programmer/chapter2"
	"github.com/fonglh/think-like-a-programmer/chapter3"
)

func chapter2Tests() {
	fmt.Println("Half Square")
	chapter2.Halfsquare()

	fmt.Println("Sideways Triangle")
	chapter2.SidewaysTriangle()

	fmt.Println("Inverted Triangle")
	chapter2.InvertedTriangle()

	fmt.Println("Diamond")
	chapter2.Diamond()

	fmt.Println("Luhn Checksum Validation")
	chapter2.ValidateLuhnChecksum()

	fmt.Println("Decode a message")
	chapter2.DecodeMessage()

	// Exercise 2-6
	fmt.Println("Convert to binary")
	chapter2.BinaryConvert()
}

func chapter3Tests() {
	fmt.Println("Find Mode")
	fmt.Println(chapter3.FindMode([]int{4, 7, 3, 8, 9, 7, 3, 9, 9, 3, 3, 10}))
}

func main() {
	fmt.Println("Chapter 3")
	chapter3Tests()
}
