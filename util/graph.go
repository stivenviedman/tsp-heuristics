package util

func GetNearestNode(node Node, candidates []Node) (Node, int, float64) {
	index := 0
	minDistance := ComputeDistance(node, candidates[0])

	for i, candidate := range candidates {
		newDist := ComputeDistance(node, candidate)

		if newDist < minDistance {
			index = i
			minDistance = newDist
		}
	}

	return candidates[index], index, minDistance
}
