package main

import (
	"fmt"
)

func main() {
	// Example usage
	s := "babad"
	result := longestPalindrome(s)
	fmt.Println(result) // Output: "bab" or "aba"
}

func longestPalindrome(s string) string {
	// iterate over s
	// start with its full length
	// check to see if it's a palindrome
	// then reduce length by 1
	// look at each substring, and check to see if palindrome
	for l := len(s); l > 0; l-- {
		fmt.Println("Checking length:", l)
		for i := 0; i+l <= len(s); i++ {
			sub := s[i : i+l]
			fmt.Println("Checking substring:", sub)
			if isPalindrome(sub) {
				return sub
			}
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	size := len(s) / 2
	str1 := s[:size]
	str2 := s[len(s)-size:]

	for i := 0; i < size; i++ {
		if str1[i] != str2[size-i-1] {
			return false
		}
	}
	return true
}
