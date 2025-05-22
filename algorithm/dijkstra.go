package algorithm

import (
	"container/heap"

	"github.com/Cosmos307/graphenalgorithmen/graph"
)

const INF = int(^uint(0) >> 1) // Maximaler int-Wert

type PQNode struct {
	Node  int
	Dist  int
	Index int // für das Heap-Interface
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

// Dijkstra-Algorithmus mit Min-Heap-Priority-Queue
func Dijkstra(g *graph.Graph, source int) (dist map[int]int, prev map[int]int) {
	dist = make(map[int]int)
	prev = make(map[int]int)
	visited := make(map[int]bool)

	for i := 0; i < g.Nodes; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQNode{Node: source, Dist: 0})

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
	return dist, prev
}
