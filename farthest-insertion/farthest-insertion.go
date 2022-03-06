package farthestinsertion

import (
	"tsp-problems/util"
)

func Run(edges []util.Edge) []util.Edge {
	var tour []util.Edge

	// Construct initial tour
	tour = append(tour, edges[0], edges[1])
	edges = util.RemoveEdge(edges, 0)
	edges = util.RemoveEdge(edges, 0)

	for len(edges) > 0 {
		// Get next edge to add
		nextEdge, index := getMaxDisEdge(tour, edges)

		// Append new edge into tour
		tour = appendNewEdge(tour, nextEdge)

		// Remove added edge from locs
		edges = util.RemoveEdge(edges, index)
	}

	return tour
}

func getMaxDisEdge(locs []util.Edge, potLocs []util.Edge) (util.Edge, int) {
	element := potLocs[0]
	index := 0
	distance := util.ComputeDistance(locs[0], potLocs[0])

	for i, canditate := range potLocs {
		for _, routeLoc := range locs {
			newDist := util.ComputeDistance(canditate, routeLoc)

			if newDist > distance {
				element = canditate
				index = i
				distance = newDist
			}
		}
	}

	return element, index
}

func appendNewEdge(locs []util.Edge, edge util.Edge) []util.Edge {
	cost := computeCost(locs[0], locs[1], edge)
	index := 0

	for i := 1; i < len(locs)-1; i++ {
		newCost := computeCost(locs[i], locs[i+1], edge)

		if newCost < cost {
			cost = newCost
			index = i
		}
	}

	return util.InsertEdge(locs, edge, index)
}

func computeCost(i util.Edge, j util.Edge, k util.Edge) float64 {
	ikDist := util.ComputeDistance(i, k)
	jkDist := util.ComputeDistance(j, k)
	ijDist := util.ComputeDistance(i, j)

	return ikDist + jkDist - ijDist
}
