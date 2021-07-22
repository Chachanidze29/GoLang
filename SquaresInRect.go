package main

import (
	"fmt"
)

func SquaresInRect(length int, width int) []int {
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

func main() {
	res := SquaresInRect(1232, 2311)
	fmt.Println(res)
}
