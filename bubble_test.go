package main

import (
	"reflect"
	"testing"
)

const ASC = 1
const DESC = -1


func TestBubbleSortInt(t *testing.T) {

	var numbers []int = []int{0, 6, 5, 3, 2, 1, 4}
	var ascSortInt = []int{0,1,2,3,4,5,6}
	var descSortInt = []int{6,5,4,3,2,1,0}
	sorted := BubbleSortInt(numbers, ASC)
	if !reflect.DeepEqual(sorted,ascSortInt) {
		t.Error("Ascending sort did not work...")
	}
	sorted = BubbleSortInt(numbers, DESC)
	if !reflect.DeepEqual(sorted,descSortInt) {
		t.Error("Descending sort did not work...")
	}
}
