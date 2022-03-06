package util

// RemoveLocation receives a slice of Point and deletes de Point corresponding to the provide index
func RemoveLocation(d []Point, i int) []Point {
	return append(d[:i], d[i+1:]...)
}

// InsertEdge inserts an edge at a specify index
func InsertEdge(d []Point, edge Point, i int) []Point {
	d = append(d[:i+1], d[i:]...)
	d[i] = edge

	return d
}
