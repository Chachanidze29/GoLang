package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	size int
	top  *Node
}

func (s *Stack) len() int {
	return s.size
}

func (s *Stack) push(val interface{}) {
	s.size++
	s.top = &Node{
		value: val,
		next:  s.top,
	}
}

func (s *Stack) pop() interface{} {
	if s.len() == 0 {
		panic("Stack is empty")
	}
	s.size--
	tmp := s.top.value
	s.top = s.top.next
	return tmp
}

func (s *Stack) printInfo() {
	curr := s.top
	for i := 0; i < s.size; i++ {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	
}
