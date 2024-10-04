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
		"rice":  5,
		"beans": 6.78,
	}

	// loop through the map above and print the values
	for k, v := range foods {
		fmt.Printf("The string key %q has a value of %v of type %T\n", k, v, v)
	}

	fmt.Println()

	// because of the nature of the type of a map[string]interface, use type assertion while looping
	for k, v := range foods {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}
}
