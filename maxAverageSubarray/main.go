package main

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(nums, k)
	println("Max average for nums", nums, "and k", k, ":", result)

	nums2 := []int{5}
	k2 := 1
	result2 := findMaxAverage(nums2, k2)
	println("Max average for nums", nums2, "and k", k2, ":", result2)
}

func findMaxAverage(nums []int, k int) float64 {
	maxVals := 0
	window := 0
	for i := 0; i < len(nums)-k; i++ {
		if i == 0 {
			for j := range nums[:i+k] {
				window += j
			}
			maxVals = window
			continue
		}
		window -= nums[i-1]
		window += nums[i+k-1]
		maxVals = max(maxVals, window)
	}
	return float64(maxVals) / float64(k)
}
