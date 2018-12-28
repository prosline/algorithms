package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("List to be sorted", numbers)
	bubbleSort(numbers)
	fmt.Println("Sorted list...", numbers)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		sweep(numbers, i)
		fmt.Println(numbers)
	}
}

func sweep(numbers []int, iterations int) {
	var N = len(numbers)
	firstIndex := 0
	secondIndex := 1
	for secondIndex < (N - iterations) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		// swapping
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}
		firstIndex++
		secondIndex++
	}
}
