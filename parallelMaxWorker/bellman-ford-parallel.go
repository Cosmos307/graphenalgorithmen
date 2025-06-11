package parallelMaxWorker

import (
	"runtime"
	"sync"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func BellmanFordParallelMaxWorker(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration, hasNegativeCycle bool) {
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	for i := range dist {
		dist[i] = algorithm.INF
		prev[i] = -1
	}
	dist[source] = 0

	numWorkers := runtime.NumCPU()
	type update struct{ to, newDist, from int }

	runtime.GC()
	start := time.Now()
	for i := 0; i < g.Nodes-1; i++ {
		tasks := make(chan int, g.Nodes)
		results := make(chan update, 1000) // moderater Buffer
		var wg sync.WaitGroup

		// Worker starten
		for w := 0; w < numWorkers; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for from := range tasks {
					for _, e := range g.Adj[from] {
						if dist[from] != algorithm.INF && dist[from]+e.Weight < dist[e.To] {
							results <- update{to: e.To, newDist: dist[from] + e.Weight, from: from}
						}
					}
				}
			}()
		}

		// Ergebnisse sammeln, während Worker laufen
		var applyWg sync.WaitGroup
		applyWg.Add(1)
		go func() {
			defer applyWg.Done()
			for upd := range results {
				if upd.newDist < dist[upd.to] {
					dist[upd.to] = upd.newDist
					prev[upd.to] = upd.from
				}
			}
		}()

		// Aufgaben verteilen
		for from := 0; from < g.Nodes; from++ {
			tasks <- from
		}
		close(tasks)
		wg.Wait()
		close(results)
		applyWg.Wait()
	}
	duration = time.Since(start)
	return dist, prev, duration, hasNegativeCycle
}
