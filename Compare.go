package main

import (
	"fmt"
	"math"
)

func Comp(array1 []int, array2 []int) bool {
	m1 := map[int]int{}
	m2 := map[int]int{}
	for i := range array2 {
		if array2[i] < 0 {
			return false
		}
	}
	for _, v := range array1 {
		m1[v] = v
	}
	for _, v := range array2 {
		m2[v] = v
	}
	for _, v := range array1 {
		if v*v != m2[v*v] {
			return false
		}
	}
	for _, v := range array2 {
		i := int(math.Sqrt(float64(v)))
		if int(i) != -m1[-i] && int(i) != m1[i] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{-9, -8, 3}
	arr2 := []int{64, 81, 9}
	fmt.Println(Comp(arr1, arr2))
}
