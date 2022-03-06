package util

import (
	"math"
	"sort"
)

// RemoveLocation receives a slice of Edges and deletes the Edge corresponding to the provide index
func RemoveEdge(d []Edge, i int) []Edge {
	return append(d[:i], d[i+1:]...)
}

// InsertEdge inserts an edge at a specified index
func InsertEdge(d []Edge, edge Edge, i int) []Edge {
	d = append(d[:i+1], d[i:]...)
	d[i] = edge

	return d
}

// ComputeDistance compute the distance between two points in 2D
func ComputeDistance(positionA Edge, positionB Edge) float64 {
	x := math.Abs(positionA.X - positionB.X)
	y := math.Abs(positionB.Y - positionA.Y)

	return math.Sqrt(x*x + y*y)
}

// ComputeTotalDistance compute the total distance of a given route
func ComputeTotalDistance(edges []Edge) float64 {
	var dis float64

	for i := 0; i < len(edges)-1; i++ {
		j := i + 1
		dis = ComputeDistance(edges[i], edges[j])
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
