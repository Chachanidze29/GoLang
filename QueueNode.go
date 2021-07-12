package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	size        int
	front, back *Node
}

func (q *Queue) len() int {
	return q.size
}

func (q *Queue) push_back(val interface{}) {
	curr := &Node{value: val, next: nil}
	if q.size == 0 {
		q.front = curr
		q.back = q.front
		q.size++
		return
	}
	q.back.next = curr
	q.back = curr
	q.size++
}

func (q *Queue) pop_front() interface{} {
	if q.len() == 0 {
		panic("Queue is empty")
	}
	tmp := q.front.value
	q.front = q.front.next
	q.size--
	return tmp
}

func (q *Queue) printInfo() {
	if q.len() == 0 {
		fmt.Println("Queue is empty")
		return
	}
	curr := q.front
	for i := 0; i < q.size; i++ {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
}

func main() {

}
