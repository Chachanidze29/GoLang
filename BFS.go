package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	size        int
	front, back *Node
}

func (q *Queue) len() int {
	return q.size
}

func (q *Queue) get_front() int {
	return q.front.value
}

func (q *Queue) push_back(val int) {
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

func (q *Queue) pop_front() int {
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

type Graph struct {
	V   int
	adj map[int][]int
}

func NewGraph() Graph {
	return Graph{
		adj: make(map[int][]int),
	}
}

func (g *Graph) addEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
}

func (g *Graph) BFS(startingNode int) {
	visited := []bool{}
	for i := 0; i < g.V; i++ {
		visited = append(visited, false)
	}
	visited[startingNode] = true
	q := Queue{}
	q.push_back(startingNode)
	for q.len() > 0 {
		curr := q.get_front()
		fmt.Print(curr, " ")
		q.pop_front()
		for _, node := range g.adj[curr] {
			if !visited[node] {
				q.push_back(node)
				visited[node] = true
			}
		}
	}
}

func main() {
	g := NewGraph()
	g.V = 4
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)
	g.BFS(2)
}
