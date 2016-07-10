package sorts

func BubbleSort(rawData []int) []int {
	// Starting from left, compare each pair in array, swapping when required
	// After each iteration of the outer for, another element at the end is in
	// Position
	for i := len(rawData) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if rawData[j] > rawData[j+1] {
				rawData[j], rawData[j+1] = rawData[j+1], rawData[j]
			}
		}
	}

	return rawData
}
