package main

import (
	"fmt"
	"log"
)

type ErrNegativeSqrt float64
// ErrNegativeSqrt Is also type of error because of overloaded Error() method
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	log.SetPrefix("Message:")
	log.SetFlags(0)
	num1, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num1)
	num2, err := Sqrt(-2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num2)
}
