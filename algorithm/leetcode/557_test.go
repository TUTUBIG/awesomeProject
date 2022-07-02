package main

import (
	"fmt"
	"testing"
)

func reverseWords(s string) string {
	ss := []byte(s)
	var j = 0
	var i = 0
	for ; i < len(ss); i++ {
		if ss[i] == ' ' {
			m := (i-j)/2 + j
			for ; j < m; j++ {
				temp := ss[i-1-j]
				ss[i-1-j] = ss[j]
				ss[j] = temp
			}
			j = i + 1
		}
	}
	m := (i-j)/2 + j
	for ; j < m; j++ {
		temp := ss[i-1-j]
		ss[i-1-j] = ss[j]
		ss[j] = temp
	}

	return string(ss)
}

func TestReverse(t *testing.T) {
	fmt.Println(reverseWords("a test"))
}
