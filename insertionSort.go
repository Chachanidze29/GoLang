package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}
