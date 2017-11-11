package main

import (
	"fmt"
	"github.com/fonglh/think-like-a-programmer/chapter2"
)

func main() {
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
}
