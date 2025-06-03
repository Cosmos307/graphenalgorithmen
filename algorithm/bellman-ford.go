package algorithm

import (
	"fmt"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func BellmanFord(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration, hasNegativeCycle bool) {
	// initialisation
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	for i := 0; i < g.Nodes; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	// algorithm
	start := time.Now()
	for i := 0; i < g.Nodes-1; i++ {
		if i%10 == 0 { // alle 10 Iterationen
			fmt.Printf("Bellman-Ford: Iteration %d/%d\n", i+1, g.Nodes-1)
		}
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
