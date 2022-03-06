package algorithms

import "tsp-problems/util"

// Run runs algorithm
func NN(locs []util.Edge) []util.Edge {
	origin := locs[0]
	toVisit := locs[1:]

	visited := []util.Edge{origin}

	for len(toVisit) > 0 {
		// get list of sorted locations
		ds := util.ComputeDistances(origin, toVisit)
		sLocations := util.SortLocation(ds)

		// add nearest location to visited list
		nearest := toVisit[sLocations[0].Index]
		visited = append(visited, nearest)

		// remove last added location from to visit list
		toVisit = util.RemoveEdge(toVisit, sLocations[0].Index)

		// reassign origin
		origin = nearest
	}

	return visited
}
