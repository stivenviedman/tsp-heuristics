package insertion

import (
	"tsp-problems/algorithms/insertion/base"
	"tsp-problems/algorithms/insertion/nodeInsertion"
	"tsp-problems/algorithms/insertion/nodeSelection"
	"tsp-problems/util"
)

func FarthestInsertion(nodes []util.Node) []util.Node {
	return base.Insertion(nodes, nodeSelection.FarthestInsertion, nodeInsertion.MinCost)
}
