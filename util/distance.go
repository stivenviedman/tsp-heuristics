package util

import "math"

// ComputeDistance compute the distance between two points in 2D
func ComputeDistance(positionA Point, positionB Point) float64 {
	x := math.Abs(positionA.X - positionB.X)
	y := math.Abs(positionB.Y - positionA.Y)

	return math.Sqrt(x*x + y*y)
}

// ComputeTotalDistance compute the total distance of a given route
func ComputeTotalDistance(locs []Point) float64 {
	var dis float64

	for i := 0; i < len(locs)-1; i++ {
		j := i + 1
		dis = ComputeDistance(locs[i], locs[j])
	}

	return dis
}

// ComputeDistances compute distances from origin to a set of destinations
func ComputeDistances(origin Point, locations []Point) []DistanceElement {
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
