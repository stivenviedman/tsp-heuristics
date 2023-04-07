package nodeSelection

import "tsp-problems/util"

func FarthestInsertion(locs []util.Node, potLocs []util.Node) (util.Node, int) {
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

func NearestInsertion(locs []util.Node, potLocs []util.Node) (util.Node, int) {
	element := potLocs[0]
	index := 0
	distance := util.ComputeDistance(locs[0], potLocs[0])

	for i, candidate := range potLocs {
		for _, routeLoc := range locs {
			newDist := util.ComputeDistance(candidate, routeLoc)

			if newDist < distance {
				element = candidate
				index = i
				distance = newDist
			}
		}
	}

	return element, index
}
