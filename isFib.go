package main

import "fmt"

func main() {
	fmt.Println(checkFib(144, 1, 1))
}

func checkFib(num, a, b int) bool {
	if num == a || num == b {
		return true
	}
	if num < a+b {
		return false
	} else if num == a+b {
		return true
	} else {
		return checkFib(num, b, a+b)
	}
}
