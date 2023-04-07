package main

import (
	"fmt"
	"tsp-problems/algorithms/insertion"
	"tsp-problems/algorithms/nn"
	"tsp-problems/constants"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("tsp-data.json")

	if algoName, ok := constants.SupportedAlgorithms[input.Algorithm]; ok {
		var route []util.Node

		switch algoName {
		case constants.SupportedAlgorithms["nn"]:
			route = nn.NN(input.Nodes)
		case constants.SupportedAlgorithms["farthest-insertion"]:
			route = insertion.FarthestInsertion(input.Nodes)
		case constants.SupportedAlgorithms["nearest-insertion"]:
			route = insertion.NearestInsertion(input.Nodes)
		case constants.SupportedAlgorithms["min-max-insertion"]:
			route = insertion.MinMaxInsertion(input.Nodes)
		}

		totDis := util.ComputeTotalDistance(route)

		fmt.Println(route)
		fmt.Println(totDis)
	} else {
		fmt.Printf("%s algorithm is not supported.\n", input.Algorithm)
	}
}
