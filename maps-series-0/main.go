package main

import (
	"fmt"
)

func main() {
	// Ways of creating a map

	// way 1: declare and then initialize
	var menu map[string]float64 // at this point value of map is nil
	fmt.Printf("menu value:%#v\n", menu)

	menu = map[string]float64{ // initializing the map
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	fmt.Printf("menu value:%#v\n", menu)

	// way 2: initialize at the point of declaration
	myMenu := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	fmt.Printf("vmyMenu alue:%#v\n", myMenu)

}
