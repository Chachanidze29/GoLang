package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(1000)
	}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[min_index], arr[i] = arr[i], arr[min_index]
	}
}
