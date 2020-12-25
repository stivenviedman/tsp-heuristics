package util

// Point indicates a Point in 2D space
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Input input data format
type Input struct {
	Points []Point `json:"points"`
}
