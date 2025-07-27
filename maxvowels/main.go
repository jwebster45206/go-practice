package main

import "fmt"

func main() {
	str := "leetcode"
	k := 3
	fmt.Println(maxVowels(str, k))
}

func maxVowels(s string, k int) int {
	// naive approach:
	// iterate over substrings within s; for each, count the vowels
	// problem with this approach: it's more complex with greater k value
	// complexity: O(n*m)

	// better solution: after checking first window, only check first and last chars
	// for each iteration
	// no need to track the substring, just max chars

	// even better solution
	// sliding window should be concerned with what left the window and what entered
	// complexity: O(n)

	// I will be tripped up a little bit on rune syntax, but algorithm should be good.
	if len(s) < k {
		return 0
	}

	maxV := 0
	lastV := 0

	for i := range len(s) - k + 1 {
		if i == 0 {
			// on first iteration, check all letters in the string
			for j := range k {
				if isVowel(s[j]) {
					lastV++
				}
			}
			maxV = lastV
			fmt.Println(maxV)
			continue
		}

		// After first window calculation:
		if isVowel(s[i-1]) { // Element sliding out
			lastV--
		}
		if isVowel(s[i+k-1]) { // Element sliding in
			lastV++
		}

		maxV = max(maxV, lastV)
		fmt.Println(maxV)
	}
	return maxV
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
