package nearestneighbor

import (
	"math"
	"sort"
	"tsp-problems/util"
)

func computeDistance(positionA util.Point, positionB util.Point) float64 {
	x := math.Abs(positionA.X - positionB.X)
	y := math.Abs(positionB.Y - positionA.Y)

	return math.Sqrt(x*x + y*y)
}

func computeDistances(origin util.Point, locations []util.Point) []dElement {
	distances := []dElement{}

	for i := 0; i < len(locations); i++ {
		distance := dElement{
			distance: computeDistance(origin, locations[i]),
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
