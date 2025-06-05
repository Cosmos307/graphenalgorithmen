package algorithm

import (
	"runtime"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func BellmanFord(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration, hasNegativeCycle bool) {
	// initialisation
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	for i := range dist {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	// algorithm
	runtime.GC()
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

	// test if the graph has negative cycles
	hasNegativeCycle = false
	// for from, adj := range g.Adj {
	// 	for _, e := range adj {
	// 		if dist[from] != INF && dist[from]+e.Weight < dist[e.To] {
	// 			hasNegativeCycle = true
	// 			break
	// 		}
	// 	}
	// }

	return dist, prev, duration, hasNegativeCycle
}
