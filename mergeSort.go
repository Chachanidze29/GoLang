package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var inf = math.MaxInt64

func main() {
	arr := make([]int, 100000)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000000)
	}
	start := time.Now()
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(time.Since(start))
}

func mergeSort(arr []int, start, end int) {
	if start < end {
		middle := (start + end) / 2
		mergeSort(arr, start, middle)
		mergeSort(arr, middle+1, end)
		merge(arr, start, middle, end)
	}
}

func merge(arr []int, left, middle, right int) {
	len1 := middle - left + 1
	len2 := right - middle
	subArr1 := make([]int, len1+1)
	subArr2 := make([]int, len2+1)
	for i := 0; i < len1; i++ {
		subArr1[i] = arr[left+i]
	}
	for i := 0; i < len2; i++ {
		subArr2[i] = arr[middle+1+i]
	}
	subArr1[len1] = inf
	subArr2[len2] = inf
	i, j := 0, 0
	for k := left; k <= right; k++ {
		if subArr1[i] <= subArr2[j] {
			arr[k] = subArr1[i]
			i++
		} else if subArr1[i] > subArr2[j] {
			arr[k] = subArr2[j]
			j++
		}
	}
}
