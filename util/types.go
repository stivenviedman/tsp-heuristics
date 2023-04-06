package util

// Node indicates the coordinates of a 2d space point
type Node struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Input input data format
type Input struct {
	Nodes     []Node `json:"nodes"`
	Algorithm string `json:"algorithm"`
}

// DistanceElement represents the distance between two Node and the index of the destination
type DistanceElement struct {
	Distance float64
	Index    int
}
