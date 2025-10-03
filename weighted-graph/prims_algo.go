package weighted_graph

import priority_queue "github.com/reruno/dsa-go/priority-queue"

func PrimsAlgo(graph *Graph) []*Edge {
	result := []*Edge{}
	markedNode := map[*Node]struct{}{}
	pq := priority_queue.NewPriorityQueue()
	n := graph.Nodes[0]
	for _, edge := range n.Adj {
		pq.Push(edge, edge.Weight)
	}
	markedNode[n] = struct{}{}
	for pq.Len() != 0 {
		item := pq.Pop()
		edge := item.Value.(*Edge)
		_, ok1 := markedNode[edge.Node1]
		_, ok2 := markedNode[edge.Node2]
		if ok1 && ok2 {
			continue
		} else if !ok1 {
			for _, e := range edge.Node1.Adj {
				pq.Push(e, e.Weight)
			}
		} else if !ok2 {
			for _, e := range edge.Node2.Adj {
				pq.Push(e, e.Weight)
			}
		}
		markedNode[edge.Node1] = struct{}{}
		markedNode[edge.Node2] = struct{}{}
		result = append(result, edge)
	}
	return result
}
