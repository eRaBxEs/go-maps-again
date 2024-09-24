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
		"egss": {
			price: 1.65,
		},
	}

	// created above is a map of a string to a struct
	fmt.Printf("menu:%#v", menu)
}
