package main

// - it will only ever be true if string lengths match; return early if they don't
// - if the strings are equal, return early.
//
// Create hash maps of characters and counts
// - since we can swap positions freely (operation 1), position never matters.
//
// Operation 2
// - Kind of confusing.
// - Since position doesn't matter though, we are swapping COUNTS effectively of a given character
//
// SO first, we make a hash map for each string.
// Then we find the characters whose counts don't match.
// If there are any characters exclusive to one of the strings, return early.
//
// NOW what to do with the diffs...
// Make another hash map, of the rune counts (number of 3's, 2's, 1's, etc)
// If the hash maps don't match, return early.
// If they do match, return true.
//

func closeStrings(word1 string, word2 string) bool {

}
