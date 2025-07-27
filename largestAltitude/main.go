package main

func largestAltitude(gain []int) int {
	// naive approach:
	// altitude and maxAltitutde vars
	// iterate over gains and calculate altitutde
	// set max to max of the 2
	// complexity: O(n)

	// would it be any faster with a hash map? not really
	// proceed with easy method

	a, maxA := 0, 0
	for _, v := range gain {
		a += v
		maxA = max(maxA, a)
	}
	return maxA
}
