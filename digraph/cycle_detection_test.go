package digraph_test

import (
	"testing"

	"github.com/reruno/dsa-go/digraph"
	"github.com/stretchr/testify/assert"
)

func TestCycleDetection1(t *testing.T) {
	n1 := &digraph.Node{Val: 1}
	n2 := &digraph.Node{Val: 2}
	n3 := &digraph.Node{Val: 3}
	n4 := &digraph.Node{Val: 4}

	g := &digraph.DiGraph{Nodes: []*digraph.Node{n1, n2, n3, n4}}
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n1)
	g.AddEdge(n3, n4)

	isCycle := digraph.DetectCycle(g)
	assert.True(t, isCycle)
}

func TestCycleDetection2(t *testing.T) {
	n1 := &digraph.Node{Val: 1}
	n2 := &digraph.Node{Val: 2}
	n3 := &digraph.Node{Val: 3}
	n4 := &digraph.Node{Val: 4}

	g := &digraph.DiGraph{Nodes: []*digraph.Node{n1, n2, n3, n4}}
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n4)

	isCycle := digraph.DetectCycle(g)
	assert.False(t, isCycle)
}

func TestCycleDetection3(t *testing.T) {
	n1 := &digraph.Node{Val: 1}
	n2 := &digraph.Node{Val: 2}
	n3 := &digraph.Node{Val: 3}
	n4 := &digraph.Node{Val: 4}

	g := &digraph.DiGraph{Nodes: []*digraph.Node{n1, n2, n3, n4}}
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n1)

	isCycle := digraph.DetectCycle(g)
	assert.True(t, isCycle)
}
