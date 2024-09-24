package main

import (
	"fmt"
)

func main() {
	// Ways of creating a map

	// way 1: declare and then initialize
	var menu map[string]float64 // at this point value of map is nil
	fmt.Printf("value:%#v\n", menu)

	menu = map[string]float64{ // initializing the map
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	fmt.Printf("value:%#v\n", menu)

}
