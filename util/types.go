package util

// Point indicates a Point in 2D space
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Input input data format
type Input struct {
	Points    []Point `json:"points"`
	Algorithm string  `json:"algorithm"`
}

// DistanceElement represents the distance between two point and the index of the destination
type DistanceElement struct {
	Distance float64
	Index    int
}
