package insertion

import "tsp-problems/util"

type nodeInsertion struct{}

func (ni *nodeInsertion) computeCost(i util.Node, j util.Node, k util.Node) float64 {
	ikDist := util.ComputeDistance(i, k)
	jkDist := util.ComputeDistance(j, k)
	ijDist := util.ComputeDistance(i, j)

	return ikDist + jkDist - ijDist
}

func (ni *nodeInsertion) MinCost(locs []util.Node, node util.Node) []util.Node {
	if len(locs) < 2 {
		return append(locs, node)
	} else {
		cost := ni.computeCost(locs[0], locs[1], node)
		index := 0

		for i := 1; i < len(locs)-1; i++ {
			newCost := ni.computeCost(locs[i], locs[i+1], node)

			if newCost < cost {
				cost = newCost
				index = i
			}
		}

		return util.InsertNode(locs, node, index)
	}
}
