package algorithm

import (
	"container/heap"
	"runtime"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/graph"
)

// max int value, used as infinity value for algorithm
const INF = int(^uint(0) >> 1)

type PQNode struct {
	Node  int
	Dist  int
	Index int
}

type PriorityQueue []*PQNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}
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

// Dijkstra algorithm for shortest paths using min-heap priority queue
func Dijkstra(g *graph.Graph, source int) (dist []int, prev []int, duration time.Duration) {
	//initialise
	dist = make([]int, g.Nodes)
	prev = make([]int, g.Nodes)
	visited := make([]bool, g.Nodes)

	for i := range dist {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQNode{Node: source, Dist: 0})

	runtime.GC()
	start := time.Now()
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*PQNode)
		if visited[u.Node] {
			continue
		}
		visited[u.Node] = true

		for _, edge := range g.Adj[u.Node] {
			v := edge.To
			weight := edge.Weight
			if dist[u.Node]+weight < dist[v] {
				dist[v] = dist[u.Node] + weight
				prev[v] = u.Node
				heap.Push(&pq, &PQNode{Node: v, Dist: dist[v]})
			}
		}
	}
	duration = time.Since(start)
	return dist, prev, duration
}
