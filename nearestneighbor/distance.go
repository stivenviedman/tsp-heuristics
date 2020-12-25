package nearestneighbor

import (
	"sort"
	"tsp-problems/util"
)

func computeDistances(origin util.Point, locations []util.Point) []dElement {
	distances := []dElement{}

	for i := 0; i < len(locations); i++ {
		distance := dElement{
			distance: util.ComputeDistance(origin, locations[i]),
			index:    i,
		}

		distances = append(distances, distance)
	}

	return distances
}

func sortLocation(distances []dElement) []dElement {
	sort.SliceStable(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	return distances
}

func removeLocation(d []util.Point, i int) []util.Point {
	return append(d[:i], d[i+1:]...)
}
