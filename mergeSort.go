package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var inf = math.MaxInt64

func main() {
	arr := make([]int, 100000000)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100000)
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
		merge2(arr, start, middle, end)
	}
}

func merge1(arr []int, start, middle, end int) {
	len1 := middle - start + 1
	len2 := end - middle
	lArr := make([]int, len1)
	rArr := make([]int, len2)
	for i := 0; i < len1; i++ {
		lArr[i] = arr[start+i]
	}
	for i := 0; i < len2; i++ {
		rArr[i] = arr[middle+1+i]
	}
	i, j := 0, 0
	k := start
	for i < len1 && j < len2 {
		if lArr[i] <= rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
		k++
	}
	for i < len1 {
		arr[k] = lArr[i]
		i++
		k++
	}
	for j < len2 {
		arr[k] = rArr[j]
		j++
		k++
	}
}

func merge2(arr []int, start, middle, end int) {
	len1 := middle - start + 1
	len2 := end - middle
	lArr := make([]int, len1+1)
	rArr := make([]int, len2+1)
	for i := 0; i < len1; i++ {
		rArr[i] = arr[start+i]
	}
	for i := 0; i < len2; i++ {
		lArr[i] = arr[middle+1+i]
	}
	lArr[len1] = inf
	rArr[len2] = inf
	i, j := 0, 0
	for k := start; k <= end; k++ {
		if lArr[i] <= rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
	}
}
