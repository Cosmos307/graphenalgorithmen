package algorithm

import "github.com/Cosmos307/graphenalgorithmen/graph"

func BellmanFord(g *graph.Graph, source int) (dist map[int]int, prev map[int]int, hasNegativeCycle bool) {
	dist = make(map[int]int)
	prev = make(map[int]int)
	for i := 0; i < g.Nodes; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	type edge struct{ from, to, weight int }
	edges := []edge{}
	for from, adj := range g.Adj {
		for _, e := range adj {
			edges = append(edges, edge{from, e.To, e.Weight})
		}
	}

	for i := 0; i < g.Nodes-1; i++ {
		for _, e := range edges {
			if dist[e.from] != INF && dist[e.from]+e.weight < dist[e.to] {
				dist[e.to] = dist[e.from] + e.weight
				prev[e.to] = e.from
			}
		}
	}

	hasNegativeCycle = false
	for _, e := range edges {
		if dist[e.from] != INF && dist[e.from]+e.weight < dist[e.to] {
			hasNegativeCycle = true
			break
		}
	}

	return dist, prev, hasNegativeCycle
}
