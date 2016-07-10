package sorts

// Helper function to merge sorted arrays
func mergeArrays(first, second, dst []int) {
	firstIndex := 0
	secondIndex := 0

	for i := range dst {
		// If no more elements in second half, add first half element
		if secondIndex == len(second) {
			dst[i] = first[firstIndex]
			firstIndex++
			// If no more elements in first half, add second half element
		} else if firstIndex == len(first) {
			dst[i] = second[secondIndex]
			secondIndex++
			// If next first half element is lower, use it
		} else if first[firstIndex] < second[secondIndex] {
			dst[i] = first[firstIndex]
			firstIndex++
			// Lastly, use a second half element
		} else {
			dst[i] = second[secondIndex]
			secondIndex++
		}
	}

}

func MergeSort(rawData []int) []int {
	// Need 2 values to sort, otherwise just return slice with single element
	if len(rawData) == 1 {
		return rawData
	}

	// Find the middle index, split, and sort each half
	halfIndex := len(rawData) / 2
	firstHalf := MergeSort(rawData[:halfIndex])
	secondHalf := MergeSort(rawData[halfIndex:])

	// Merge the sorted values into a created working array
	workingArray := make([]int, len(rawData))
	mergeArrays(firstHalf, secondHalf, workingArray)

	return workingArray
}

// This doesn't perform better than a synchronous version, but its interesting to code
func ConcurrentMergeSort(rawData []int) []int {
	//Chan to receive the sorted and merged array
	done := make(chan []int)

	go asyncSplitAndMerge(rawData, done)

	sorted := <-done
	return sorted
}

func asyncSplitAndMerge(rawData []int, done chan []int) {
	//We need 2 values to sort... otherwise it's already sorted (with itself)
	if len(rawData) == 1 {
		done <- rawData
		return
	}

	//Two channels to receive sorted slices
	firstHalfDoneChan := make(chan []int)
	secondHalfDoneChan := make(chan []int)

	//Find the middle index, split, and async get back a sorted array on chan
	halfIndex := len(rawData) / 2
	go asyncSplitAndMerge(rawData[:halfIndex], firstHalfDoneChan)
	go asyncSplitAndMerge(rawData[halfIndex:], secondHalfDoneChan)

	firstHalf := <-firstHalfDoneChan
	secondHalf := <-secondHalfDoneChan

	//merge the sorted values into a created working array
	workingArray := make([]int, len(rawData))
	mergeArrays(firstHalf, secondHalf, workingArray)

	//Write out the merged array
	done <- workingArray
}
