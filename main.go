package main

import (
	"fmt"
	"time"

	"github.com/Cosmos307/graphenalgorithmen/algorithm"
	"github.com/Cosmos307/graphenalgorithmen/graph"
	"github.com/Cosmos307/graphenalgorithmen/parallel"
)

func main() {
	sizes := []int{10000}
	densities := []float64{0.01, 0.1, 0.5, 0.9}
	runs := 5

	fmt.Println("start")

	for _, n := range sizes {
		for _, d := range densities {
			g := graph.NewRandomGraph(n, d, 1, 10)
			fmt.Printf("\n--- Graph: n=%d, density=%.2f ---\n", n, d)

			avgDijkstra, avgDijkstraPara, avgBellman, avgBellmanPara := execAllAlgoAvg(g, 0, runs)
			fmt.Printf("Dijkstra (avg):           %v\n", avgDijkstra)
			fmt.Printf("Dijkstra Parallel (avg):  %v\n", avgDijkstraPara)
			fmt.Printf("Bellman-Ford (avg):       %v\n", avgBellman)
			fmt.Printf("Bellman-Ford Parallel (avg): %v\n", avgBellmanPara)
		}
	}

	// Lineare Graphen
	for _, n := range sizes {
		g := graph.NewLinearGraph(n, 1, 10)
		fmt.Printf("\n--- Linearer Graph: n=%d ---\n", n)

		avgDijkstra, avgDijkstraPara, avgBellman, avgBellmanPara := execAllAlgoAvg(g, 0, runs)
		fmt.Printf("Dijkstra (avg):           %v\n", avgDijkstra)
		fmt.Printf("Dijkstra Parallel (avg):  %v\n", avgDijkstraPara)
		fmt.Printf("Bellman-Ford (avg):       %v\n", avgBellman)
		fmt.Printf("Bellman-Ford Parallel (avg): %v\n", avgBellmanPara)
	}
}

func execAllAlgoAvg(g *graph.Graph, start int, runs int) (avgDijkstra, avgDijkstraPara, avgBellman, avgBellmanPara time.Duration) {
	var sumDijkstra, sumDijkstraPara, sumBellman, sumBellmanPara time.Duration

	for i := 0; i < runs; i++ {
		_, _, durDijkstra := algorithm.Dijkstra(g, start)
		sumDijkstra += durDijkstra

		_, _, durDijkstraPara := parallel.DijkstraParallel(g, start)
		sumDijkstraPara += durDijkstraPara

		_, _, durBellman, _ := algorithm.BellmanFord(g, start)
		sumBellman += durBellman

		_, _, durBellmanPara, _ := parallel.BellmanFordParallel(g, start)
		sumBellmanPara += durBellmanPara
	}

	avgDijkstra = sumDijkstra / time.Duration(runs)
	avgDijkstraPara = sumDijkstraPara / time.Duration(runs)
	avgBellman = sumBellman / time.Duration(runs)
	avgBellmanPara = sumBellmanPara / time.Duration(runs)
	return
}
