package insertion

import "tsp-problems/util"

type nodeSelection struct{}

type localNearestPoint struct {
	index    int
	distance float64
}

func (nodeSelection *nodeSelection) FarthestInsertion(locs []util.Node, potLocs []util.Node) (util.Node, int) {
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

func (nodeSelection *nodeSelection) NearestInsertion(locs []util.Node, potLocs []util.Node) (util.Node, int) {
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

// Node whose minimal distance to a tour node is maximal
func (nodeSelection *nodeSelection) MinMax(locs []util.Node, potLocs []util.Node) (util.Node, int) {
	index := 0
	minDistances := map[int]localNearestPoint{}

	for i, candidate := range potLocs {
		_, _, minDist := util.GetNearestNode(candidate, locs)
		minDistances[i] = localNearestPoint{i, minDist}
	}

	for _, localNearest := range minDistances {
		if localNearest.distance > minDistances[index].distance {
			index = localNearest.index
		}
	}

	return potLocs[index], index
}
