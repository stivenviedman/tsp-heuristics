package insertion

import "tsp-problems/util"

func NearestInsertion(nodes []util.Node) []util.Node {
	nodeInsertion := nodeInsertion{}
	nodeSelection := nodeSelection{}

	return baseInsertion(nodes, nodeSelection.NearestInsertion, nodeInsertion.MinCost)
}
