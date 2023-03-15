package math

// FindMax returns the maximum integer value in the input slice of integers.
// If the slice is empty, returns 0.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}
