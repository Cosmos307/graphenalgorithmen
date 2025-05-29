package main

import (
	"fmt"
	"log"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
)

func main() {
	g := graph.NewRandomGraph(10, 0.3, 1, 10)
	fmt.Println("--- Generierter Graph ---")
	g.Print()

	// algorithm
	dist, prev, duration := algorithm.Dijkstra(g, 0)

	distBF, prevBF, durationBF, negCycle := algorithm.BellmanFord(g, 0)
	if negCycle {
		log.Fatal("Negativer Zyklus gefunden!")
	}

	//Print Dijkstra algorithm results
	fmt.Println("\n--- Dijkstra (Startknoten 0) ---")
	for node := 0; node < g.Nodes; node++ {
		fmt.Printf("Knoten %d: Distanz = %d, Vorgänger = %d\n", node, dist[node], prev[node])
	}
	fmt.Printf("Dijkstra Duration: %d", duration)

	//Print Bellman-Ford algorithm results
	fmt.Println("\n--- Bellman-Ford (Startknoten 0) ---")
	for node := 0; node < g.Nodes; node++ {
		fmt.Printf("Knoten %d: Distanz = %d, Vorgänger = %d\n", node, distBF[node], prevBF[node])
	}
	fmt.Printf("Bellman-Ford: %d", durationBF)
}
