package week01

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	var s1, s2 = "anagram", "nagaram"
	if isAnagram(s1, s2) {
		t.Log("true")
		return
	}
	t.Log("false")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := [26]int{}
	ts := [26]int{}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		ss[index]++
	}
	fmt.Print(ss)
	for i := 0; i < len(t); i++ {
		index := t[i] - 'a'
		ts[index]++
	}
	fmt.Print(ss)
	return ss == ts
}
