package parallel

import (
	"container/heap"
	"runtime"
	"sync"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
)

type PQNode struct {
	Node  int
	Dist  int
	Index int
}

type PriorityQueue []*PQNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Dist < pq[j].Dist }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*PQNode)
	node.Index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// Parallel version of Dijkstra algorithm: neighbour-relaxation parallelized
func DijkstraParallel(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration) {
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	visited := make([]bool, g.Nodes)
	for i := range dist {
		dist[i] = algorithm.INF
		prev[i] = -1
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQNode{Node: source, Dist: 0})

	var mu sync.Mutex

	runtime.GC()
	start := time.Now()
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*PQNode)
		if visited[u.Node] {
			continue
		}
		visited[u.Node] = true

		var wg sync.WaitGroup
		for _, edge := range g.Adj[u.Node] {
			wg.Add(1)
			go func(edge graph.Edge) {
				defer wg.Done()
				v := edge.To
				weight := edge.Weight
				mu.Lock()
				if dist[u.Node]+weight < dist[v] {
					dist[v] = dist[u.Node] + weight
					prev[v] = u.Node
					heap.Push(&pq, &PQNode{Node: v, Dist: dist[v]})
				}
				mu.Unlock()
			}(edge)
		}
		wg.Wait()
	}
	duration = time.Since(start)
	return dist, prev, duration
}
