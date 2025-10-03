package weighted_graph

import "fmt"

func NewNode(val int) *Node {
	return &Node{
		Val: val,
		Adj: []*Edge{},
	}
}

type Node struct {
	Val int
	Adj []*Edge
}

func NewEdge(n1, n2 *Node, weight int) *Edge {
	return &Edge{
		Node1:  n1,
		Node2:  n2,
		Weight: weight,
	}
}

type Edge struct {
	Node1  *Node
	Node2  *Node
	Weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d", e.Weight)
}

func NewGraph(nodes ...*Node) *Graph {
	return &Graph{
		Nodes: nodes,
	}
}

type Graph struct {
	Nodes []*Node
}

func (w *Graph) AddEdge(n1 *Node, n2 *Node, edge *Edge) {
	n1.Adj = append(n1.Adj, edge)
	n2.Adj = append(n2.Adj, edge)
}
