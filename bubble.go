package main

import "fmt"

const ASC = 1
const DESC = -1

func main() {
	var numbers []int = []int{0, 6, 5, 3, 2, 1, 4}
	bubbleSortInt(numbers, ASC)
	fmt.Println("List to be sorted in Ascending Order", numbers)
	bubbleSortInt(numbers, DESC)
	fmt.Println("List to be sorted in Descending Order", numbers)
}

func bubbleSortInt(numbers []int, order int) {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		if !sweep(numbers, i, order) {
			return
		}
		fmt.Println(numbers)
	}
}

func sweep(numbers []int, iterations int, order int) bool {
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
