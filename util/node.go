package util

import (
	"math"
	"sort"
)

// RemoveNode receives a slice of nodes and deletes the node corresponding to the provide index
func RemoveNode(d []Node, i int) []Node {
	return append(d[:i], d[i+1:]...)
}

// InsertNode inserts an node at a specified index
func InsertNode(d []Node, node Node, i int) []Node {
	d = append(d[:i+1], d[i:]...)
	d[i] = node

	return d
}

// ComputeDistance compute the distance between two points in 2D
func ComputeDistance(positionA Node, positionB Node) float64 {
	x := math.Abs(positionA.X - positionB.X)
	y := math.Abs(positionB.Y - positionA.Y)

	return math.Sqrt(x*x + y*y)
}

// ComputeTotalDistance compute the total distance of a given route
func ComputeTotalDistance(nodes []Node) float64 {
	var dis float64

	for i := 0; i < len(nodes)-1; i++ {
		j := i + 1
		dis = ComputeDistance(nodes[i], nodes[j])
	}

	return dis
}

// ComputeDistances compute distances from origin to a set of destinations
func ComputeDistances(origin Node, locations []Node) []DistanceElement {
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
