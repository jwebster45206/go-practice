package main

func main() {
	// Example usage
	n := 10 // Example value for n
	result := guessNumber(n)
	println(result) // Output will depend on the implementation of guess function
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// guessNumber uses textbook binary search to guess a number between 1 and n
// complexity: O(log n) -> very low complexity, which is good.
func guessNumber(n int) int {
	floor := 1
	ceil := n
	for {
		g := floor + (ceil-floor)/2 // TODO check, but I think this will round down
		r := guess(g)
		switch r {
		case -1:
			// guess was too high; look in lower half next
			ceil = g
		case 1:
			// guess was too low; look in upper half next
			if floor == g {
				// special case where we're close to the result
				// assumes that the guess is rounded down
				floor = g + 1
				continue
			}
			floor = g
		case 0:
			return g
		}
	}
}

func guess(num int) int {
	// This is a mock implementation of the guess function.
	// In a real scenario, this function would be provided by the system.
	const pickedNumber = 6 // Example picked number for testing
	if num < pickedNumber {
		return 1 // num is lower than the picked number
	} else if num > pickedNumber {
		return -1 // num is higher than the picked number
	}
	return 0 // num is equal to the picked number
}
