package main

import "fmt"

func main() {
	arr1 := []int{1, 0, 1, 0, 0, 1}
	arr2 := []int{1, 1, 1, 1, 0, 1}
	res := addition(arr1, arr2, 6)
	fmt.Println(res)
}

func addition(a, b []int, n int) []int {
	res := make([]int, n+1)
	carry := 0
	for i := n; i > 0; i-- {
		sum := a[i-1] + b[i-1] + carry
		if sum == 1 || sum == 3 {
			res[i] = 1
		} else {
			res[i] = 0
		}
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
	}
	res[0] = carry
	return res
}
