package nearestneighbor

import "tsp-problems/util"

// Run runs algorithm
func Run(locs []util.Point) []util.Point {
	origin := locs[0]
	toVisit := locs[1:]

	visited := []util.Point{origin}

	for len(toVisit) > 0 {
		// get list of sorted locations
		ds := util.ComputeDistances(origin, toVisit)
		sLocations := sortLocation(ds)

		// add nearest location to visited list
		nearest := toVisit[sLocations[0].Index]
		visited = append(visited, nearest)

		// remove last added location from to visit list
		toVisit = removeLocation(toVisit, sLocations[0].Index)

		// reassign origin
		origin = nearest
	}

	return visited
}
