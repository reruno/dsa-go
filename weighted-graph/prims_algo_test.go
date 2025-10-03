package weighted_graph_test

import (
	"testing"

	weighted_graph "github.com/reruno/dsa-go/weighted-graph"
	"github.com/stretchr/testify/assert"
)

func prepGraph() *weighted_graph.Graph {
	n0 := weighted_graph.NewNode(0)
	n1 := weighted_graph.NewNode(1)
	n2 := weighted_graph.NewNode(2)
	n3 := weighted_graph.NewNode(3)
	n4 := weighted_graph.NewNode(4)
	n5 := weighted_graph.NewNode(5)
	n6 := weighted_graph.NewNode(6)
	n7 := weighted_graph.NewNode(7)
	edge07 := weighted_graph.NewEdge(n0, n7, 1)
	edge23 := weighted_graph.NewEdge(n2, n3, 2)
	edge17 := weighted_graph.NewEdge(n1, n7, 3)
	edge02 := weighted_graph.NewEdge(n0, n2, 4)
	edge57 := weighted_graph.NewEdge(n5, n7, 5)
	edge13 := weighted_graph.NewEdge(n1, n3, 6)
	edge15 := weighted_graph.NewEdge(n1, n5, 7)
	edge27 := weighted_graph.NewEdge(n2, n7, 8)
	edge45 := weighted_graph.NewEdge(n4, n5, 9)
	edge12 := weighted_graph.NewEdge(n1, n2, 10)
	edge47 := weighted_graph.NewEdge(n4, n7, 11)
	edge04 := weighted_graph.NewEdge(n0, n4, 12)
	edge62 := weighted_graph.NewEdge(n6, n2, 13)
	edge36 := weighted_graph.NewEdge(n3, n6, 14)
	edge60 := weighted_graph.NewEdge(n6, n0, 15)
	edge64 := weighted_graph.NewEdge(n6, n4, 16)
	graph := weighted_graph.NewGraph(n0, n1, n2, n3, n4, n5, n6, n7)
	graph.AddEdge(n0, n7, edge07)
	graph.AddEdge(n0, n7, edge07)
	graph.AddEdge(n2, n3, edge23)
	graph.AddEdge(n1, n7, edge17)
	graph.AddEdge(n0, n2, edge02)
	graph.AddEdge(n5, n7, edge57)
	graph.AddEdge(n1, n3, edge13)
	graph.AddEdge(n1, n5, edge15)
	graph.AddEdge(n2, n7, edge27)
	graph.AddEdge(n4, n5, edge45)
	graph.AddEdge(n1, n2, edge12)
	graph.AddEdge(n4, n7, edge47)
	graph.AddEdge(n0, n4, edge04)
	graph.AddEdge(n6, n2, edge62)
	graph.AddEdge(n3, n6, edge36)
	graph.AddEdge(n6, n0, edge60)
	graph.AddEdge(n6, n4, edge64)
	return graph
}

func TestPrimsAlgo1(t *testing.T) {
	graph := prepGraph()
	path := weighted_graph.PrimsAlgo(graph)
	sum := 0
	for _, edge := range path {
		sum += edge.Weight
	}
	assert.Equal(t, 37, sum)
}
