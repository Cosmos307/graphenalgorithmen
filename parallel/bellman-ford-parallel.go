package parallel

import (
	"runtime"
	"sync"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func BellmanFordParallel(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration, hasNegativeCycle bool) {
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	for i := 0; i < g.Nodes; i++ {
		dist[i] = algorithm.INF
		prev[i] = -1
	}
	dist[source] = 0

	runtime.GC()
	start := time.Now()
	for i := 0; i < g.Nodes-1; i++ {
		var wg sync.WaitGroup
		var mu sync.Mutex
		updates := make([][3]int, 0)
		for from := 0; from < g.Nodes; from++ {
			wg.Add(1)
			adj := g.Adj[from]
			go func(from int, adj []graph.Edge) {
				defer wg.Done()
				for _, e := range adj {
					if dist[from] != algorithm.INF && dist[from]+e.Weight < dist[e.To] {
						mu.Lock()
						updates = append(updates, [3]int{e.To, dist[from] + e.Weight, from})
						mu.Unlock()
					}
				}
			}(from, adj)
		}
		wg.Wait()
		for _, upd := range updates {
			to, newDist, from := upd[0], upd[1], upd[2]
			if newDist < dist[to] {
				dist[to] = newDist
				prev[to] = from
			}
		}
	}
	duration = time.Since(start)

	// Negative cycle detection
	hasNegativeCycle = false
	// for from, adj := range g.Adj {
	// 	for _, e := range adj {
	// 		if dist[from] != algorithm.INF && dist[from]+e.Weight < dist[e.To] {
	// 			hasNegativeCycle = true
	// 			break
	// 		}
	// 	}
	// }
	return dist, prev, duration, hasNegativeCycle
}
