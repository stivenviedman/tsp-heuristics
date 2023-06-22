package insertion

import "tsp-problems/util"

func baseInsertion(
	nodes []util.Node,
	selectNode func(locs []util.Node, potLocs []util.Node) (util.Node, int),
	insertNode func(locs []util.Node, node util.Node) []util.Node,
) []util.Node {
	var tour []util.Node

	// Construct initial tour
	tour = append(tour, nodes[0])
	nodes = util.RemoveNode(nodes, 0)

	for len(nodes) > 0 {
		// Select next node to add to the partial tour
		nextNode, index := selectNode(tour, nodes)

		// Insert node in the partial tour
		tour = insertNode(tour, nextNode)

		// Remove added node from locs
		nodes = util.RemoveNode(nodes, index)
	}

	return tour
}
