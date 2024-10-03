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
}
