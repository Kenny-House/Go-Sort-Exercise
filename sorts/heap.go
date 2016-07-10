package sorts

func HeapSort(rawData []int) []int {
	// Output slice referencing same backing array
	outputSlice := rawData[0:]
	limit := len(outputSlice)

	// Initial heapification, ensures that all child nodes can be safely assumed valid
	heapify(rawData)

	for i := 0; i < limit; i++ {
		//Move max element from heap onto output, set last element from raw data as root
		outputSlice[len(rawData)-1], rawData[0] = rawData[0], rawData[len(rawData)-1]
		//Resize raw data to exclude last element
		rawData = rawData[:len(rawData)-1]
		//Sift down the new root element
		siftDown(rawData, 0, len(rawData))
	}

	return rawData
}

// Take an array and create a valid heap by starting at the highest index parent
// And creating a valid heap via sifting the parent down the child heaps, then
// Progressing up through parent nodes until 0 index is reached
func heapify(rawData []int) {
	// Start to sift down at the last parent node and move up
	root := getParentIndex(len(rawData) - 1)
	for root >= 0 {
		siftDown(rawData, root, len(rawData))
		root--
	}
}

// Repair heap rooted at root, assumes heaps rooted at children are valid
// Heap array, parent to start on, last index in the heap array
func siftDown(heap []int, root, limit int) {
	lChild, rChild := getChildIndices(root)

	//If there are no children, just return
	if lChild >= limit {
		return
	}

	//Assume root index to be greatest
	// log.Printf("Assume index %d max", root)
	max := root

	// If parent < left child, left child should be considered the max
	if heap[max] < heap[lChild] {
		// log.Printf("Assume lChild %d max", lChild)
		max = lChild
	}

	// If right child exists is greater than max (which may be root or left child)
	// Set right as max
	if rChild < limit && heap[max] < heap[rChild] {
		// log.Printf("rChild %d is max", rChild)
		max = rChild
	}

	// If root is not max, swap max and root
	// Continue sifting down
	if max != root {
		heap[max], heap[root] = heap[root], heap[max]
		siftDown(heap, max, limit)
	}

}

func getChildIndices(i int) (l, r int) {
	return (2 * i) + 1, (2 * i) + 2
}

func getParentIndex(c int) int {
	return (c - 1) / 2
}
