package arrays

func Sum(arr []int) int {
	sum := 0

	// _ is index & number is value
	for _, number := range arr {
		sum += number
	}

	return sum
}
