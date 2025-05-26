package main

import (
	"fmt"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func main() {
	g := graph.NewRandomGraph(10, 0.3, 1, 10)
	fmt.Println("--- Generierter Graph ---")
	g.Print()

	dist, prev := algorithm.Dijkstra(g, 0)
	fmt.Println("\n--- Dijkstra (Startknoten 0) ---")
	for node := 0; node < g.Nodes; node++ {
		fmt.Printf("Knoten %d: Distanz = %d, Vorgänger = %d\n", node, dist[node], prev[node])
	}

	distBF, prevBF, negCycle := algorithm.BellmanFord(g, 0)
	fmt.Println("\n--- Bellman-Ford (Startknoten 0) ---")
	if negCycle {
		fmt.Println("Warnung: Negativer Zyklus gefunden!")
	} else {
		for node := 0; node < g.Nodes; node++ {
			fmt.Printf("Knoten %d: Distanz = %d, Vorgänger = %d\n", node, distBF[node], prevBF[node])
		}
	}
}
