package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func printNode(n *Node) {
	fmt.Println("Value :", n.value)
	if n.left != nil {
		fmt.Println("Left node value :", n.left.value)
	}
	if n.right != nil {
		fmt.Println("right node value :", n.right.value)
	}
	fmt.Println()
}

func main() {
	var N int
	fmt.Scan(&N)
	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scan(&val, &indexLeft, &indexRight)
		nodes[i].value = val
		if indexLeft >= 0 && indexLeft <= N {
			nodes[i].left = &nodes[indexLeft]
		}
		if indexRight >= 0 && indexRight <= N {
			nodes[i].right = &nodes[indexRight]
		}
	}
	for _, node := range nodes {
		printNode(&node)
	}
}
