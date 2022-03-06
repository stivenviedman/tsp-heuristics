package util

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
