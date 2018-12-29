package main

import "strings"

/*
func main() {
	var numbers []int = []int{0, 6, 5, 3, 2, 1, 4}
	bubbleSortInt(numbers, ASC)
	fmt.Println("List to be sorted in Ascending Order", numbers)
	bubbleSortInt(numbers, DESC)
	fmt.Println("List to be sorted in Descending Order", numbers)
}
*/

// Sorting Slice of Numbers
func BubbleSortInt(numbers []int, order int) []int {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		if !sweepInt(numbers, i, order) {
			return numbers
		}
	}
	return numbers
}

func sweepInt(numbers []int, iterations int, order int) bool {
	var N = len(numbers)
	firstIndex := 0
	secondIndex := 1
	swapped := false
	for secondIndex < (N - iterations) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		// swapping
		if order > 0 {
			if firstNumber > secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				swapped = true
			}

		} else {
			if firstNumber < secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				swapped = true
			}

		}
		firstIndex++
		secondIndex++
	}
	return swapped
}

// Sorting Slice of Strings
func BubbleSortString(names []string, order int) []string {
	var N int = len(names)
	for i := 0; i < N; i++ {
		if !sweepString(names, i, order) {
			return names
		}
	}
	return names
}

func sweepString(names []string, iterations int, order int) bool {
	var N = len(names)
	firstIndex := 0
	secondIndex := 1
	swapped := false
	for secondIndex < (N - iterations) {
		var firstNumber string = names[firstIndex]
		var secondNumber string = names[secondIndex]
		// swapping
		if order > 0 {
			if strings.ToLower(firstNumber) > strings.ToLower(secondNumber) {
				names[firstIndex] = secondNumber
				names[secondIndex] = firstNumber
				swapped = true
			}

		} else {
			if strings.ToLower(firstNumber) < strings.ToLower(secondNumber) {
				names[firstIndex] = secondNumber
				names[secondIndex] = firstNumber
				swapped = true
			}

		}
		firstIndex++
		secondIndex++
	}
	return swapped
}
