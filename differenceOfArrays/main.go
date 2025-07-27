package main

import "fmt"

const isInNums1 = 1
const isInNums2 = 2
const nonExclusive = 0

func main() {
	// Example usage:
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	result := findDifference(nums1, nums2)
	fmt.Println(result) // Output: [[1, 3], [4, 6]]
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	// return 2 slices
	// each should contain the items that are exclusive from the original lists

	// if we loop through nums1, and check every value in nums2 as we go:
	// complexity will be O(n*m)

	// we could build a map containing
	// - item value, and list presence
	// complexity more like O(n+m) which is better
	m := make(map[int]int)

	// then loop through each slice
	for _, v := range nums1 {
		mapCompare(v, isInNums1, m)
	}
	for _, v := range nums2 {
		mapCompare(v, isInNums2, m)
	}

	// build final lists
	f := make([][]int, 2)
	for k, v := range m {
		if v == isInNums1 {
			f[0] = append(f[0], k)
			continue
		}
		if v == isInNums2 {
			f[1] = append(f[1], k)
		}
	}
	return f
}

func mapCompare(v int, listStatus int, m map[int]int) {
	// is the item value already in the map?
	// if for this array, continue
	// if for the other array, mark it non-exclusive
	if status, ok := m[v]; ok {
		if status == listStatus {
			return
		}
		m[v] = nonExclusive
		return
	}
	m[v] = listStatus
}
