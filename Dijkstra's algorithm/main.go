package main

import (
	"fmt"
	"math"
)

// IsElemInSlice is func linear search element in slice.
func IsElemInSlice(elem string, sl []string) bool {
	for _, v := range sl {
		if v == elem {
			return true
		}
	}
	return false
}

// SearchLowCostNode find the lowest cost among nods
func SearchLowCostNode(m map[string]float64, pro []string) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""
	for i, val := range m {
		if val < lowestCost && !IsElemInSlice(i, pro) {
			lowestCost = val
			lowestCostNode = i
		}
	}
	return lowestCostNode
}

// Program use Dijkstra`s algorithm for search min way from start to final in DAG(Directed acyclic graph
//without negative weights of edge)
func main() {
	// Init graph of ways from start to final
	graph := make(map[string]map[string]float64)
	graph["start"] = map[string]float64{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = map[string]float64{}
	graph["a"]["final"] = 1
	graph["b"] = map[string]float64{}
	graph["b"]["a"] = 3
	graph["b"]["final"] = 5

	costs := map[string]float64{}
	costs["a"] = 6
	costs["b"] = 2
	costs["final"] = math.Inf(1)

	// Map parents needed to build min way from start to final
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["final"] = ""

	// Slice processed needed to mark processed nodes
	processed := make([]string, 0)

	node := SearchLowCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		for i, val := range graph[node] {
			newCost := cost + val
			if newCost < costs[i] {
				costs[i] = newCost
				parents[i] = node
			}
		}
		processed = append(processed, node)
		node = SearchLowCostNode(costs, processed)
	}

	fmt.Printf("Min cost from start to final %v\n", costs["final"])
	way := "final"
	for i := "final"; i != "start"; {
		i = parents[i]
		way += "-" + i
	}
	fmt.Println("Way is", way)
}
