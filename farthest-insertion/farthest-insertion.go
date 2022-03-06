package farthestinsertion

import (
	"tsp-problems/util"
)

func Run(locs []util.Point) []util.Point {
	var tour []util.Point

	// Construct initial tour
	tour = append(tour, locs[0], locs[1])
	locs = util.RemoveLocation(locs, 0)
	locs = util.RemoveLocation(locs, 0)

	for len(locs) > 0 {
		// Get next edge to add
		nextEdge, index := getMaximumDisPoint(tour, locs)

		// Append new edge into tour
		tour = appendNewEdge(tour, nextEdge)

		// Remove added edge from locs
		locs = util.RemoveLocation(locs, index)
	}

	return tour
}

func getMaximumDisPoint(locs []util.Point, potLocs []util.Point) (util.Point, int) {
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

func appendNewEdge(locs []util.Point, edge util.Point) []util.Point {
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

func computeCost(i util.Point, j util.Point, k util.Point) float64 {
	ikDist := util.ComputeDistance(i, k)
	jkDist := util.ComputeDistance(j, k)
	ijDist := util.ComputeDistance(i, j)

	return ikDist + jkDist - ijDist
}
