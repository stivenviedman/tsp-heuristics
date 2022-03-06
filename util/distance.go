package util

import (
	"math"
	"sort"
)

// ComputeDistance compute the distance between two points in 2D
func ComputeDistance(positionA Edge, positionB Edge) float64 {
	x := math.Abs(positionA.X - positionB.X)
	y := math.Abs(positionB.Y - positionA.Y)

	return math.Sqrt(x*x + y*y)
}

// ComputeTotalDistance compute the total distance of a given route
func ComputeTotalDistance(locs []Edge) float64 {
	var dis float64

	for i := 0; i < len(locs)-1; i++ {
		j := i + 1
		dis = ComputeDistance(locs[i], locs[j])
	}

	return dis
}

// ComputeDistances compute distances from origin to a set of destinations
func ComputeDistances(origin Edge, locations []Edge) []DistanceElement {
	distances := []DistanceElement{}

	for i := 0; i < len(locations); i++ {
		distance := DistanceElement{
			Distance: ComputeDistance(origin, locations[i]),
			Index:    i,
		}

		distances = append(distances, distance)
	}

	return distances
}

// SortLocation sorts location from closest to farthest
func SortLocation(distances []DistanceElement) []DistanceElement {
	sort.SliceStable(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return distances
}
