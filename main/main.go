package main

import (
	"fmt"
	"tsp-problems/constants"
	"tsp-problems/nearestneighbor"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("tsp-data.json")

	if util.Contains(constants.SupportedAlgorithms, input.Algorithm) {
		route := nearestneighbor.Run(input.Points)
		totDis := util.ComputeTotalDistance(route)

		fmt.Println(nearestneighbor.Run(input.Points))
		fmt.Println(totDis)
	} else {
		fmt.Printf("%s algorithm is not supported.\n", input.Algorithm)
	}
}
