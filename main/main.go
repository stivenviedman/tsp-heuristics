package main

import (
	"fmt"
	"tsp-problems/constants"
	farthestinsertion "tsp-problems/farthest-insertion"
	"tsp-problems/nearestneighbor"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("tsp-data.json")

	if util.Contains(constants.SupportedAlgorithms, input.Algorithm) {
		var route []util.Point

		switch input.Algorithm {
		// "nn"
		case constants.SupportedAlgorithms[0]:
			route = nearestneighbor.Run(input.Points)
		// farthest insertion
		case constants.SupportedAlgorithms[1]:
			route = farthestinsertion.Run(input.Points)
		}

		totDis := util.ComputeTotalDistance(route)

		fmt.Println(route)
		fmt.Println(totDis)
	} else {
		fmt.Printf("%s algorithm is not supported.\n", input.Algorithm)
	}
}
