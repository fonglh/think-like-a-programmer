package chapter3

// According to the problem, the numbers are in the range 1-10
func FindMode(input []int) int {
	var histogram [10]int

	// Create histogram
	for _, value := range input {
		histogram[value-1] += 1
	}

	return getIndexofHighestCount(histogram[:])
}

func getIndexofHighestCount(input []int) int {
	highestCount := 0
	idxOfHighest := 0

	for index, count := range input {
		if count > highestCount {
			highestCount = count
			idxOfHighest = index + 1
		}
	}

	return idxOfHighest
}
