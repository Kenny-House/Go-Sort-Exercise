package sorts

func InsertionSort(rawData []int) []int {
	// Start at second element and go to last element
	for i := 1; i < len(rawData); i++ {
		// If the element at the current index (i) is less than the preceding element,
		// Swap the elements. Continue this from right to left, swapping until the
		// Element is in place
		for j := i; j > 0 && rawData[j-1] > rawData[j]; j-- {
			rawData[j-1], rawData[j] = rawData[j], rawData[j-1]
		}
	}

	return rawData
}
