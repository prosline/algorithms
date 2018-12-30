package main

// Search for value in the slice and returns the index if found. Otherwise, -1
func BinarySearch(numbers []int, numberToFind int) int {

	// Example, []numbers{1,2,3,4,5,6,7,8,9,10,12,13,14,15,16,17,18,19,20}
	var lo int = 0
	var hi int = len(numbers) - 1

	for lo <= hi {

		var eleMid int = lo + (hi-lo)/2
		var midValue int = numbers[eleMid]

		if midValue == numberToFind {
			return eleMid

		} else if midValue > numberToFind {
			// Use the first half of the slice
			hi = eleMid - 1

		} else { // midValue < numberToFind
			// Use the first half of the slice
			lo = eleMid + 1
		}
	}

	return -1

}
