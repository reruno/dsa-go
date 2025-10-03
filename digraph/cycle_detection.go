package digraph

func DetectCycle(g *DiGraph) bool {
	marked := map[int]struct{}{}
	restack := map[int]struct{}{}

	for _, n := range g.Nodes {
		if _, ok := marked[n.Val]; !ok {
			isCycle := detectCycleDfs(n, marked, restack)
			if isCycle {
				return isCycle
			}
		}
	}
	return false
}

func detectCycleDfs(n *Node, marked map[int]struct{}, restack map[int]struct{}) bool {
	marked[n.Val] = struct{}{}
	if len(n.Adj) == 0 {
		return false
	}

	_, ok := marked[n.Val]
	_, okRestack := restack[n.Val]
	if ok && okRestack {
		return true
	}
	restack[n.Val] = struct{}{}
	for _, adj := range n.Adj {
		isCycle := detectCycleDfs(adj, marked, restack)
		if isCycle {
			return true
		}
	}
	delete(restack, n.Val)
	return false
}
