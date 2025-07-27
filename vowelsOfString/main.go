package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example usage
	s := "leetcode"
	result := reverseVowels(s)
	println(result) // Output: "eoo"
}

func reverseVowels(s string) string {
	const vowels = "aeiouAEIOU"
	vowelIndices := []int{}
	vowelValues := []rune{}

	// get the indexes and values of all vowels
	runes := []rune(s)
	for i, r := range runes {
		if strings.Contains(vowels, string(r)) {
			vowelIndices = append(vowelIndices, i)
			vowelValues = append(vowelValues, r)
		}
	}

	fmt.Println("Vowel Indices:", vowelIndices)

	// iterate over indexes, replacing with values from the end minus current index
	for k := range vowelIndices {
		fmt.Println(k)
		if k > (len(vowelIndices)-1)/2 {
			break
		}
		fmt.Println(k)
		j := len(vowelIndices) - 1 - k
		a := runes[vowelIndices[k]]
		b := runes[vowelIndices[j]]
		runes[vowelIndices[k]] = b
		runes[vowelIndices[j]] = a
	}
	return string(runes)
}
