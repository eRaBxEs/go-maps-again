package main

import (
	"fmt"
)

// imagine a struct exists
type menuItem struct {
	price float64
}

func main() {
	// creating a map using a map literal
	menu := map[string]menuItem{
		"beans": {
			price: 1.65,
		},
	}

	// created above is a map of a string to a struct
	fmt.Printf("menu:%#v", menu)

	// The cannot assign to struct field map error
	// lets try to change the price of beans
	// menu["beans"].price = 0.25 will give us the cannot assign to struct field map error, instead do the below
	beans := menu["beans"]
	beans.price = 0.25    // assign a value to the price field
	menu["beans"] = beans // update the price

}
