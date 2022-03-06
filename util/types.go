package util

// Edge indicates a Edge in 2D space
type Edge struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Input input data format
type Input struct {
	Edges     []Edge `json:"edges"`
	Algorithm string `json:"algorithm"`
}

// DistanceElement represents the distance between two Edges and the index of the destination
type DistanceElement struct {
	Distance float64
	Index    int
}
