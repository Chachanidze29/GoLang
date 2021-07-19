package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type pNode *Node

type Deque struct {
	size  int
	front pNode
	back  pNode
}

func (d *Deque) len() int {
	return d.size
}

func (d *Deque) push_front(value int) {
	if d.size == 0 {
		d.front = &Node{value: value, next: nil, prev: nil}
		d.back = d.front
		d.size++
		return
	}
	d.front = &Node{value: value, next: d.front, prev: nil}
	d.size++
}

func (d *Deque) pop_front() int {
	if d.size == 0 {
		panic("Deque is empty")
	}
	tmp := d.front.value
	d.front = d.front.next
	d.size--
	return tmp
}

func (d *Deque) push_back(value int) {
	curr := &Node{value: value}
	if d.size == 0 {
		d.back = curr
		d.front = d.back
		d.size++
		return
	}
	d.size++
	curr.prev = d.back
	d.back.next = curr
	d.back = curr
}

func (d *Deque) pop_back() int {
	if d.size == 0 {
		panic("Deque is empty")
	}
	d.size--
	tmp := d.back.value
	d.back = d.back.prev
	if d.back == nil {
		d.front = nil
		return tmp
	}
	d.back.next = nil
	return tmp
}

func (d *Deque) printInfo() {
	curr := d.front
	for i := 0; i < d.size; i++ {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
}

func main() {
	d := Deque{}
	for i := 0; i < 10; i++ {
		d.push_back(i)
	}
	d.printInfo()
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(d.pop_front(), " ")
	}
}
