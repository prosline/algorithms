package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	size := 20
	slice := createSlice(size)
	fmt.Println("Unsorted slice....", slice)
	selectionSort(slice)
	fmt.Println("Sorted slice....", slice)

}

func createSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var midIndex = 0
		for j := i; j < n; j++ {
			if items[j] < items[midIndex] {
				midIndex = j
			}
		}
		items[i], items[midIndex] = items[midIndex], items[i]
	}
}
