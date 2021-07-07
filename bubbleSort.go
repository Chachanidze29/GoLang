package main

import "fmt"

func main() {
	var numbers []int = []int{1, 5, 2, 3, 6, 7, 1, 2}
	fmt.Println("Our list of numbers", numbers)
	for i := 0; i < len(numbers); i++ {
		if !sweep(numbers, i) {
			break
		}
	}
	fmt.Println("Our list of numbers (sorted)", numbers)
}

func sweep(numbers []int, prevPasses int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false
	for secondIndex < N-prevPasses {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
