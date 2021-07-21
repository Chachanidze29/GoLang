package main

import "fmt"

func Min(arr []int) int {
	m := map[int]int{}
	for i, v := range arr {
		m[i] = v
	}
	for i := 0; i < len(m); i++ {
		if i != m[i] {
			return i
		}
	}
	return len(m)
}

func main() {
	arr := []int{-100, -12332, -1313}
	fmt.Println(Min(arr))
}
