package main

import (
	"fmt"
)

func main() {
	// create a map using map literal
	menu := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	// making a list of only the keys, which is the menu here
	for dish := range menu {
		fmt.Println("Please remember to get", dish)
	}

	// if we want only the second value of the map, we use a blank identifier to replace the key
	var total float64
	for _, price := range menu {
		total += price
	}

	fmt.Printf("An everything breakfast will cost you %.02f", total)
}
