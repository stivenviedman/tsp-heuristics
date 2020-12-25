package main

import (
	"fmt"
	"tsp-problems/nearestneighbor"
	"tsp-problems/util"
)

func main() {
	input := util.ReadTxt("test.json")
	fmt.Println(nearestneighbor.Run(input.Points))
}
