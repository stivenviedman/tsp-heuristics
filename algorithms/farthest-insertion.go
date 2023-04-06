package algorithms

import "tsp-problems/util"

func FarthestInsertion(nodes []util.Node) []util.Node {
	return insertion(nodes, getFarthestNode, insertMinCost)
}

func insertion(
	nodes []util.Node,
	selectNode func(locs []util.Node, potLocs []util.Node) (util.Node, int),
	insertNode func(locs []util.Node, node util.Node) []util.Node,
) []util.Node {
	var tour []util.Node

	// Construct initial tour
	tour = append(tour, nodes[0], nodes[1])
	nodes = util.RemoveNode(nodes, 0)
	nodes = util.RemoveNode(nodes, 0)

	for len(nodes) > 0 {
		// Get next RemoveNode to add
		nextNode, index := selectNode(tour, nodes)

		// Append new RemoveNode into tour
		tour = insertNode(tour, nextNode)

		// Remove added node from locs
		nodes = util.RemoveNode(nodes, index)
	}

	return tour
}

func getFarthestNode(locs []util.Node, potLocs []util.Node) (util.Node, int) {
	element := potLocs[0]
	index := 0
	distance := util.ComputeDistance(locs[0], potLocs[0])

	for i, candidate := range potLocs {
		for _, routeLoc := range locs {
			newDist := util.ComputeDistance(candidate, routeLoc)

			if newDist > distance {
				element = candidate
				index = i
				distance = newDist
			}
		}
	}

	return element, index
}

func insertMinCost(locs []util.Node, node util.Node) []util.Node {
	cost := computeCost(locs[0], locs[1], node)
	index := 0

	for i := 1; i < len(locs)-1; i++ {
		newCost := computeCost(locs[i], locs[i+1], node)

		if newCost < cost {
			cost = newCost
			index = i
		}
	}

	return util.InsertNode(locs, node, index)
}

func computeCost(i util.Node, j util.Node, k util.Node) float64 {
	ikDist := util.ComputeDistance(i, k)
	jkDist := util.ComputeDistance(j, k)
	ijDist := util.ComputeDistance(i, j)

	return ikDist + jkDist - ijDist
}
