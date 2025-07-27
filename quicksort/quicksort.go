package main

import "time"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quicksort(left), pivot), quicksort(right)...)
}

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

func main() {

	nums := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42, 13, 27, 38, 45, 67, 89, 23, 56, 78,
		91, 33, 55, 77, 99, 21, 43, 65, 87, 19, 41, 63, 85, 17, 39, 61, 83, 15, 37, 59,
		81, 14, 36, 58, 80, 16, 40, 62, 84, 18, 35, 57, 79, 20, 44, 66, 88, 24, 46, 68,
		92, 26, 48, 70, 94, 28, 52, 74, 96, 30, 54, 76, 98, 32, 53, 75, 97, 31, 51, 73,
		95, 29, 49, 71, 93, 27, 47, 69, 91, 25, 45, 67, 89, 23, 43, 65, 87, 21, 41, 63,
		85, 19, 39, 61, 83, 17, 37, 59, 81, 15, 35, 57, 79, 13, 33, 55, 77, 11, 31, 53,
		75, 9, 29, 51, 73, 7, 27, 49, 71, 5, 25, 47, 69, 3, 23, 45, 67, 1, 21, 43, 65,
		2, 22, 44, 66, 4, 24, 46, 68, 6, 26, 48, 70, 8, 28, 50, 72, 10, 30, 52, 74, 12,
		32, 54, 76, 14, 34, 56, 78, 16, 36, 58, 80, 18, 38, 60, 82, 20, 40, 62, 84, 100,
		101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
		117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130}

	start := time.Now()
	_ = quicksort(nums)
	duration := time.Since(start).Microseconds()
	println("Execution time (quicksort):", duration, "microseconds")

	nums = []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42, 13, 27, 38, 45, 67, 89, 23, 56, 78,
		91, 33, 55, 77, 99, 21, 43, 65, 87, 19, 41, 63, 85, 17, 39, 61, 83, 15, 37, 59,
		81, 14, 36, 58, 80, 16, 40, 62, 84, 18, 35, 57, 79, 20, 44, 66, 88, 24, 46, 68,
		92, 26, 48, 70, 94, 28, 52, 74, 96, 30, 54, 76, 98, 32, 53, 75, 97, 31, 51, 73,
		95, 29, 49, 71, 93, 27, 47, 69, 91, 25, 45, 67, 89, 23, 43, 65, 87, 21, 41, 63,
		85, 19, 39, 61, 83, 17, 37, 59, 81, 15, 35, 57, 79, 13, 33, 55, 77, 11, 31, 53,
		75, 9, 29, 51, 73, 7, 27, 49, 71, 5, 25, 47, 69, 3, 23, 45, 67, 1, 21, 43, 65,
		2, 22, 44, 66, 4, 24, 46, 68, 6, 26, 48, 70, 8, 28, 50, 72, 10, 30, 52, 74, 12,
		32, 54, 76, 14, 34, 56, 78, 16, 36, 58, 80, 18, 38, 60, 82, 20, 40, 62, 84, 100,
		101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
		117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130}

	start = time.Now()
	_ = selectionSort(nums)
	duration = time.Since(start).Microseconds()
	println("Execution time (selection sort):", duration, "microseconds")

	//println("Sorted array:", sorted)
}
