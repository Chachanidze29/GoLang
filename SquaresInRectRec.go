package main

import (
	"fmt"
)

func SquaresInRect(length, width int) []int {
	res := []int{}
	if length == width {
		return nil
	}
	for true {
		if length > width {
			length -= width
			res = append(res, width)
		} else if length == width {
			res = append(res, length)
			break
		} else {
			width -= length
			res = append(res, length)
		}
	}
	return res
}

func SquaresInRectRec(length, width int, arr []int) []int {
	if length == width {
		return append(arr, length)
	}
	if length > width {
		arr = append(arr, width)
		return SquaresInRectRec(length-width, width, arr)
	}
	if length < width {
		arr = append(arr, length)
		return SquaresInRectRec(length, width-length, arr)
	}
	return arr
}

func main() {
	arr := []int{}
	res := SquaresInRectRec(1232, 2311, arr)
	fmt.Println(res)
}
