package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "tsp-problems/nearestneighbor"
)

// ReadTxt reads file and return slice of locations
func ReadTxt(file string) Input {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %v, %v", file, err)
	}

	var data Input

	err = json.Unmarshal(f, &data)
	if err != nil {
		fmt.Printf("JSON input not valid %v", err)
	}

	return data
}
