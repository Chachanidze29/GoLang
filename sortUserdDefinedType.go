package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Person struct {
	name string
	age  int
}

type byAge []Person

func (arrP byAge) Len() int {
	return len(arrP)
}

func (arrP byAge) Less(i, j int) bool {
	return arrP[i].age < arrP[j].age
}

func (arrP byAge) Swap(i, j int) {
	arrP[i], arrP[j] = arrP[j], arrP[i]
}

func main() {
	rand.Seed(time.Now().Unix())
	arr := []Person{}
	for i := 0; i < 100; i++ {
		arr = append(arr, Person{"a", rand.Int() % 1000000})
	}
	start := time.Now()
	sort.Sort(byAge(arr))
	finish := time.Now()
	elapsed := finish.Sub(start)
	fmt.Println(elapsed)
	b := byAge{Person{"AVTO", 20}}
	fmt.Println(b)
}
