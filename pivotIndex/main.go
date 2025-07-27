package main

import "fmt"

func main() {
	// Example usage
	nums := []int{1, 7, 3, 6, 5, 6}
	result := pivotIndex(nums)
	println(result) // Output: 3
}

func pivotIndex(nums []int) int {
	// for each index, we will need to know the sum of all values to left and all values to right
	// there can be more than one.
	// return first occurrence of one (leftmost)

	// on first run, calc. all values to right.
	// on subsequent runs, add the -1 element to leftVals and subtract the 0 element from rightVals
	// whenever leftVals and rightVals are equal, return the index

	vals1 := 0
	vals2 := 0
	for i, v := range nums {
		if i == 0 {
			for j := 1; j < len(nums); j++ {
				vals2 += nums[j]
			}
		} else {
			vals1 += nums[i-1]
			vals2 -= v
		}
		fmt.Println("i:", i, "vals1:", vals1, "vals2:", vals2)
		if vals1 == vals2 {
			return i
		}
	}
	return -1
}
