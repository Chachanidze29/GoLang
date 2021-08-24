package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(linearSearchRec(arr, len(arr), 1))
}

func linearSearchRec(arr []int, N int, val int) bool {
	if N == 0 {
		return false
	}
	if val == arr[N-1] {
		return true
	}
	return linearSearchRec(arr, N-1, val)
}
