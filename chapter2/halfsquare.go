package chapter2

import "fmt"

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
		for j := i; j > 0; j-- {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
