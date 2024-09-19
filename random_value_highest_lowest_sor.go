package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input array
	numbers := []int{3, 6, 1, 7, 5}
	
	// Find minimum and maximum values
	min := numbers[0]
	max := numbers[0]
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Sort the array
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)

	// Output the results
	fmt.Printf("Minimum = %d\n", min)
	fmt.Printf("Maximum = %d\n", max)
	fmt.Printf("Sorted Array = %v\n", sortedNumbers)
}
