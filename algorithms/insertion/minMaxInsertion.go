package insertion

import "tsp-problems/util"

func MinMaxInsertion(nodes []util.Node) []util.Node {
	nodeInsertion := nodeInsertion{}
	nodeSelection := nodeSelection{}

	return baseInsertion(nodes, nodeSelection.MinMax, nodeInsertion.MinCost)
}
