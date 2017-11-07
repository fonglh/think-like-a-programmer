package chapter2

import (
	"fmt"
	"math"
)

// Prints '#' symbols, as given by rowLength.
func printRow(rowLength int) {
	for i := rowLength; i > 0; i-- {
		fmt.Printf("#")
	}
	fmt.Printf("\n")
}

func Halfsquare() {
	/*
		Print a pattern of hash symbols shaped like half of a perfect
		5 x 5 square.
		#####
		####
		###
		##
		#
	*/
	for i := 5; i > 0; i-- {
		printRow(i)
	}
}

func Sidewaystriangle() {
	/*
		Print a pattern of hash symbols shaped like a sideways triangle.
		#
		##
		###
		####
		###
		##
		#
	*/
	for row := 1; row < 8; row++ {
		numHash := (int)(4 - math.Abs(float64(4-row)))
		printRow(numHash)
	}
}
