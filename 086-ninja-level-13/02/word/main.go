// Package word provides functions to count words in a string.
package word

import "strings"

// Count counts the number of words in a string s.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}

// UseCount returns a map with the count of each word in the string s.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range xs {
		m[word]++
	}
	return m
}