package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo :", i)
		time.Sleep(time.Duration(3 * time.Microsecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar :", i)
		time.Sleep(time.Duration(20 * time.Microsecond))
	}
	wg.Done()
}

func main() {

}
