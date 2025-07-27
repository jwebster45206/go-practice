package main

import "fmt"

func main() {
	// This is a placeholder for the main function.
	// You can add your application logic here.

	println(mergeAlternately("abc", "pqr"))
	println(mergeAlternately("ab", "pqrs"))
	println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	w := ""
	// start a counter
	// for each index, add the char from word1, if it exists
	// add the char from word2, if it exists
	// if neither exists, break
	i := 0
	for {
		switch {
		case len(word1) > i:
			w += string(word1[i])
			fallthrough
		case len(word2) > i:
			fmt.Println(len(word2), i)
			w += string(word2[i])
		case len(word1) < i && len(word2) < i:
			return w
		}
		i++
	}
}
