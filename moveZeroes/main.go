package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	i := 0
	numsCount := len(nums)
	zeroCount := 0
	for i < numsCount-zeroCount {
		if nums[i] != 0 {
			i++
			zeroCount++
			continue
		}
		for j := i; j < numsCount; j++ {
			if j == numsCount-1 {
				nums[j] = 0
				break
			}
			nums[j] = nums[j+1]
		}
	}
}
