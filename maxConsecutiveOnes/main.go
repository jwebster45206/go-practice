package main

// Given a binary array nums and an integer k, return the maximum number of
// consecutive 1's in the array if you can flip at most k 0's.

func main() {
	// Example usage
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	result := longestOnes(nums, k)
	println(result) // Output: 6
}

func longestOnes(nums []int, k int) int {
	maxOnes := 0
	for i := range nums {
		if len(nums)-i <= maxOnes {
			break
		}
		setOnes := 0
		setExceptions := 0
		for j := i; j < len(nums); j++ { // range nums[i:]
			if nums[j] == 0 {
				if setExceptions >= k {
					break
				}
				setExceptions++
			}
			setOnes++
		}
		maxOnes = max(maxOnes, setOnes)
	}
	return maxOnes
}
