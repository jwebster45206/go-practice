package main

import "strings"

func main() {
	// Example usage
	str1 := "ABCABC"
	str2 := "ABC"
	result := gcdOfStrings(str1, str2)
	println(result) // Output: "ABC"
}

func gcdOfStrings(str1 string, str2 string) string {
	// iterate over character combinations in string 1, checking
	// for existence in string 2
	// start with all of string 1, and work with progressively smaller slices of the string
	// return as soon as a match is found
	for z := len(str1); z > 0; z++ {
		if len(str2) < z {
			continue
		}
		for i := 0; i+z <= len(str1); i++ {
			s := str1[i : i+z]
			if strings.Contains(str2, s) {
				return s
			}
		}
	}
	return ""
}
