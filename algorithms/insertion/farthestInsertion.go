package insertion

import "tsp-problems/util"

func FarthestInsertion(nodes []util.Node) []util.Node {
	nodeInsertion := nodeInsertion{}
	nodeSelection := nodeSelection{}

	return baseInsertion(nodes, nodeSelection.FarthestInsertion, nodeInsertion.MinCost)
}
