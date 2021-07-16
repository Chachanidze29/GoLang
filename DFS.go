package main

import "fmt"

type Graph struct {
	V   int
	adj map[int][]int
}

func NewGraph(vNum int) Graph {
	return Graph{
		adj: make(map[int][]int),
		V:   vNum,
	}
}

func (g *Graph) addEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
}

func (g *Graph) dfsviz(v int, visited []bool) {
	visited[v] = true
	fmt.Print(v, " ")
	for _, node := range g.adj[v] {
		if !visited[node] {
			visited[node] = true
			g.dfsviz(node, visited)
		}
	}
}

func (g *Graph) DFS(v int) {
	visited := []bool{}
	for i := 0; i < g.V; i++ {
		visited = append(visited, false)
	}
	g.dfsviz(v, visited)
	fmt.Println()
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.dfsviz(i, visited)
			fmt.Println()
		}
	}
}

func main() {
	g := NewGraph(6)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 4)
	g.addEdge(1, 3)
	g.addEdge(4, 5)
	g.addEdge(1, 0)
	g.addEdge(2, 0)
	g.addEdge(4, 0)
	g.addEdge(5, 4)
	g.addEdge(3, 1)
	g.DFS(0)
}

