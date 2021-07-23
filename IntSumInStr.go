package main

import (
	"fmt"
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	temp := "0"

	Sum := 0

	for _, ch := range strng {

		if ch >= 48 && ch <= 57 {
			temp += string(ch)
		} else {
			num, _ := strconv.Atoi(temp)
			Sum += num
			temp = "0"
		}
	}
	num, _ := strconv.Atoi(temp)
	return Sum + num
}

func main() {
	word := "12abc20yz68"
	fmt.Println(SumOfIntegersInString(word))
}
