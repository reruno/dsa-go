package digraph

import "fmt"

type Node struct {
	Val int
	Adj []*Node
}

type DiGraph struct {
	Nodes []*Node
}

func (g *DiGraph) AddEdge(n1, n2 *Node) {
	n1.Adj = append(n1.Adj, n2)
}

func (g *DiGraph) Print() {
	fmt.Println("Undirected Graph:")
	for _, node := range g.Nodes {
		fmt.Printf("Node %d -> ", node.Val)
		for _, adj := range node.Adj {
			fmt.Printf("%d ", adj.Val)
		}
		fmt.Println()
	}
}
