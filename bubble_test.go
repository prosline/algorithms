package main

import (
	"reflect"
	"testing"
)

const ASC = 1
const DESC = -1


func TestBubbleSortInt(t *testing.T) {

	// Testing Slice of integers
	var numbers []int = []int{0, 6, 5, 3, 2, 1, 4}
	var ascSortInt = []int{0,1,2,3,4,5,6}
	var descSortInt = []int{6,5,4,3,2,1,0}
	sorted := BubbleSortInt(numbers, ASC)
	if !reflect.DeepEqual(sorted,ascSortInt) {
		t.Error("Ascending sort [Slice of Integer] did not work...")
	}
	sorted = BubbleSortInt(numbers, DESC)
	if !reflect.DeepEqual(sorted,descSortInt) {
		t.Error("Descending sort [Slice of Integer] did not work...")
	}

	//----------------------------------------------------------------------
	// Testing Slice of strings
	var names []string = []string{"John", "Paul", "Anna", "Richard", "Carl"}
	var ascSortString = []string{"Anna","Carl", "John", "Paul", "Richard"}
	var descSortString = []string{"Richard", "Paul", "John", "Carl", "Anna"}
	sortedString := BubbleSortString(names, ASC)
	if !reflect.DeepEqual(sortedString,ascSortString) {
		t.Error("Ascending sort [Slice of String] did not work...")
	}
	sortedString = BubbleSortString(names, DESC)
	if !reflect.DeepEqual(sortedString,descSortString) {
		t.Error("Descending sort [Slice of String] did not work...")
	}
}
