package chapter3

// According to the problem, the numbers are in the range 1-10
func FindMode(input []int) int {
	var counts [10]int

	for _, value := range input {
		counts[value-1] += 1
	}

	highestCount := 0
	currMode := 0
	for index, count := range counts {
		if count > highestCount {
			highestCount = count
			currMode = index + 1
		}
	}
	return currMode
}
