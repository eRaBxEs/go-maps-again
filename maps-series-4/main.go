package main

import (
	"fmt"
	"sort"
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

	// ordering map keys; ordering maps in a certain manner
	var dishes []string // use a slice
	for dish := range menu {
		dishes = append(dishes, dish)
	}

	fmt.Printf("\ndishes arrangement before sorting: %#v\n", dishes)
	// Sorting alphabetically
	sort.Strings(dishes)
	fmt.Printf("dishes arrangement after sorting: %#v\n", dishes)
	// now using the ordered slice, print the map key and the map value
	fmt.Println("The menu in alphabetical order:")
	for _, dish := range dishes {
		fmt.Println(dish, menu[dish])
	}

}
