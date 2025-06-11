package parallelMaxWorker

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

func DijkstraParallelMaxWorker(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration) {
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

	numWorkers := runtime.NumCPU()
	type task struct {
		u    int
		edge graph.Edge
	}

	runtime.GC()
	start := time.Now()
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*PQNode)
		if visited[u.Node] {
			continue
		}
		visited[u.Node] = true

		tasks := make(chan task, len(g.Adj[u.Node]))
		var wg sync.WaitGroup
		var mu sync.Mutex

		// Worker starten
		for w := 0; w < numWorkers; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for t := range tasks {
					v := t.edge.To
					weight := t.edge.Weight
					mu.Lock()
					if dist[t.u]+weight < dist[v] {
						dist[v] = dist[t.u] + weight
						prev[v] = t.u
						heap.Push(&pq, &PQNode{Node: v, Dist: dist[v]})
					}
					mu.Unlock()
				}
			}()
		}

		// Aufgaben verteilen
		for _, edge := range g.Adj[u.Node] {
			tasks <- task{u: u.Node, edge: edge}
		}
		close(tasks)
		wg.Wait()
	}
	duration = time.Since(start)
	return dist, prev, duration
}
