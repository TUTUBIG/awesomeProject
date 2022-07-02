package main

import (
	"testing"
)

func isPalindrome1(s string) bool {
	// remove non-alphanumeric characters
	ss := make([]byte, 0, len(s))
	for i := range s {
		// lowercase character minus upercase character equals to 26
		if s[i] >= 'A' && s[i] <= 'Z' {
			ss = append(ss, s[i]+32)
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			ss = append(ss, s[i])
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			ss = append(ss, s[i])
			continue
		}
	}

	middle := len(ss) / 2
	for i := 0; i < middle; i++ {
		if ss[i] != ss[len(ss)-i-1] {
			return false
		}
	}
	return true
}
func TestIsPalindrome(t *testing.T) {
	isPalindrome1("0P")
}
