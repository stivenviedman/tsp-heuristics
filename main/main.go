package main

import (
	"fmt"
	"tsp-problems/nearestneighbor"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("test.json")
	route := nearestneighbor.Run(input.Points)
	totDis := util.ComputeTotalDistance(route)

	fmt.Println(nearestneighbor.Run(input.Points))
	fmt.Println(totDis)
}
