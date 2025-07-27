package main

// Selection Sort implementation in Go
// Complexity: O(n^2) -> not the best, but easy to implement and understand
//
// Different from the book, but more efficient because it avoids
// memory reallocation and multiple swaps.
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ { // Position to fill with next smallest
		minIndex := i                // Assume current position has minimum
		for j := i + 1; j < n; j++ { // Search remaining unsorted portion
			if arr[j] < arr[minIndex] {
				minIndex = j // Found smaller element
			}
		}
		// Swap the minimum with position i
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}
