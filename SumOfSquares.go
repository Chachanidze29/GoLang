package main

import "fmt"

func main() {
	fmt.Println(sumOfSquare(100))
}

func sumOfSquare(n int) int {
	if n == 1 {
		return 1
	}
	return sumOfSquare(n-1) + n*n
}
