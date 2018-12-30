package main

import "testing"

func TestBinarySearch(t *testing.T) {

	var numbers []int =  []int{1,2,3,4,5,6,7,8,9,10,12,13,14,15,16,17,18,19,20}

	value := BinarySearch(numbers, 9)

	if !(value == 8){
		t.Error("Value of 8 not found. It found instead ",value)
	}
}
