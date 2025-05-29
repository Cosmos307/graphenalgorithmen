package algorithm

import (
	"time"

	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func BellmanFord(g *graph.Graph, source int) (dist map[int]int, prev map[int]int, duration time.Duration, hasNegativeCycle bool) {
	// initialisation
	dist = make(map[int]int)
	prev = make(map[int]int)
	for i := 0; i < g.Nodes; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	start := time.Now()
	for i := 0; i < g.Nodes-1; i++ {
		for from, adj := range g.Adj {
			for _, e := range adj {
				if dist[from] != INF && dist[from]+e.Weight < dist[e.To] {
					dist[e.To] = dist[from] + e.Weight
					prev[e.To] = from
				}
			}
		}
	}
	duration = time.Since(start)

	hasNegativeCycle = false
	for from, adj := range g.Adj {
		for _, e := range adj {
			if dist[from] != INF && dist[from]+e.Weight < dist[e.To] {
				hasNegativeCycle = true
				break
			}
		}
	}

	return dist, prev, duration, hasNegativeCycle
}
