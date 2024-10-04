package main

import (
	"fmt"
)

func main() {
	// create a map of a string interface using map literal
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{
			"chicken",
			1.75,
		},
		"steak": true,
	}

	// loop through the map above and print the values
	for k, v := range foods {
		fmt.Printf("The string key %q has a value of %v of type %T\n", k, v, v)
	}
}
