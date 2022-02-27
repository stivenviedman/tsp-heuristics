package nearestneighbor

import (
	"sort"
	"tsp-problems/util"
)

func sortLocation(distances []util.DistanceElement) []util.DistanceElement {
	sort.SliceStable(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return distances
}

func removeLocation(d []util.Point, i int) []util.Point {
	return append(d[:i], d[i+1:]...)
}
