package main

import "strings"

func main() {
	// Example usage
	s := "hello world"
	result := reverseWords(s)
	println(result) // Output: "olleh dlrow"
}

func reverseWords(s string) string {
	// divide the string by spaces
	// words := strings.Split(s, " ")
	words := strings.Fields(s)

	// reverse the order of the words
	// make a slice of strings of the correct length
	words2 := make([]string, len(words))
	for k, v := range words {
		words2[len(words2)-1-k] = v
	}
	return strings.Join(words2, " ")

	// Important things to remember:
	// strings.Split
	// strings.Fields
	// strings.Join
	// iterating over a range
	// making a slice
}
