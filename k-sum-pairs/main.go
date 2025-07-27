package main

import (
	"log"
)

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	result := maxOperations(nums, k)
	log.Printf("Max operations for nums %v and k %d: %d\n", nums, k, result)

	nums2 := []int{3, 1, 3, 4, 3}
	k2 := 6
	result2 := maxOperations(nums2, k2)
	log.Printf("Max operations for nums %v and k %d: %d\n", nums2, k2, result2)
}

func maxOperations(nums []int, k int) int {

	vals := make(map[int]int)
	for i := range nums {
		key := nums[i]
		if _, exists := vals[key]; exists {
			vals[key]++
			continue
		}
		vals[key] = 1
	}

	matches := 0
	for i := range vals {
		if i >= k {
			continue
		}
		if _, exists := vals[k-i]; exists {
			matches += min(vals[i], vals[k-i])
		}
	}
	return matches / 2

}
