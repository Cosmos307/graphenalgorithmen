package graph

import (
	"fmt"
	"math/rand"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Nodes int
	Adj   map[int][]Edge
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.Adj[from] = append(g.Adj[from], Edge{To: to, Weight: weight})
}

func NewRandomGraph(nodes int, density float64, minWeight int, maxWeight int) *Graph {
	g := &Graph{
		Nodes: nodes,
		Adj:   make(map[int][]Edge),
	}

	weightRange := maxWeight - minWeight + 1

	for from := 0; from < nodes; from++ {
		for to := 0; to < nodes; to++ {
			if from != to && rand.Float64() < density {
				weight := rand.Intn(weightRange) + minWeight
				g.AddEdge(from, to, weight)
			}

		}
	}
	return g
}

func (g *Graph) Print() {
	for from, edges := range g.Adj {
		for _, edge := range edges {
			fmt.Printf("%d -> %d (Gewicht: %d)\n", from, edge.To, edge.Weight)
		}
	}
}
