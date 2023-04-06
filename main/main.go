package main

import (
	"fmt"
	"tsp-problems/algorithms"
	"tsp-problems/constants"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("tsp-data.json")

	if util.Contains(constants.SupportedAlgorithms, input.Algorithm) {
		var route []util.Node

		switch input.Algorithm {
		// "nn"
		case constants.SupportedAlgorithms[0]:
			route = algorithms.NN(input.Nodes)
		// farthest insertion
		case constants.SupportedAlgorithms[1]:
			route = algorithms.FarthestInsertion(input.Nodes)
		}

		totDis := util.ComputeTotalDistance(route)

		fmt.Println(route)
		fmt.Println(totDis)
	} else {
		fmt.Printf("%s algorithm is not supported.\n", input.Algorithm)
	}
}
