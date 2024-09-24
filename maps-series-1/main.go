package main

import (
	"fmt"
)

func main() {
	// creating a map using map literal
	menu := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	// to retrieve the price of eggs; use the name of the map and the key to access
	eggsPrice := menu["eggs"]
	fmt.Printf("eggsPrice: %v\n", eggsPrice)

	// to change the price of eggs say to 1.50; use the name of the map and then the key and pass in the desired value
	menu["eggs"] = 1.50
	fmt.Printf("menu:%#v", menu)

}
