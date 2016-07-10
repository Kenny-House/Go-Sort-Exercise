package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Kenny-House/Go-Sorting/sorts"
)

const ELEMENTS int = 1000
const LIMIT int = ELEMENTS * 10
const TEST_RUNS int = 10

type sorterFunc func([]int) []int

func main() {
	log.Printf("Elements: %v", ELEMENTS)
	log.Printf("Test runs per sort function: %v", TEST_RUNS)

	// bubble := runTest(sorts.BubbleSort)
	// log.Printf("Bubble Sort: %d", bubble)

	// insertion := runTest(sorts.InsertionSort)
	// log.Printf("Insertion Sort: %d", insertion)

	// merge := runTest(sorts.MergeSort)
	// log.Printf("Merge Sort: %d", merge)
	//
	heap := runTest(sorts.HeapSort)
	log.Printf("Heap Sort: %d", heap)

}

func runTest(sf sorterFunc) (elapsed int64) {
	rawData := make([]int, ELEMENTS)

	for i := 0; i < TEST_RUNS; i++ {
		fillRandomData(rawData)
		start := time.Now()
		sf(rawData)
		elapsed += time.Since(start).Nanoseconds()
	}

	elapsed = elapsed / int64(TEST_RUNS)

	return elapsed
}

func fillRandomData(rawData []int) {
	rand.Seed(time.Now().UnixNano())

	for i := range rawData {
		rawData[i] = rand.Int() % LIMIT
	}
}
