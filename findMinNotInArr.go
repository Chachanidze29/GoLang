package main

import "fmt"

func isInArr(arr []int, val int) bool {
	for i := range arr {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func Min(arr []int) int {
	for i := 0; i < len(arr)+1; i++ {
		if !isInArr(arr, i) {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{-1, -2, -3, -5, 0, 1}
	fmt.Println(Min(arr))
}
