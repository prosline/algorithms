package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {

	// Testing Slices of integer
	var arg1 []int = []int{1,2,3,2}
	var arg2 []int = []int{9,1,4,4}
	ok := ValidAnagram(arg1, arg2)
	if !ok {
		t.Error("Slices are not anagrams")
	}
}
