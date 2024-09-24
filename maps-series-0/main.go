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

	fmt.Printf("myMenu alue:%#v\n", myMenu)

	// way 3: using the make function to create a map
	anotherMenu := make(map[string]float64) // Nb:at this point map is non-nil which signifies that it is just empty and not at the default zero type of nil
	fmt.Printf("anotherMenu value:%#v\n", anotherMenu)

	// way 4: Creating a ready to use empty map wihout using make function
	yourMenu := map[string]float64{}
	fmt.Printf("yourMenu value:%#v\n", yourMenu)

	// initializing both anotherMenu and yourMenu
	anotherMenu["eggs"], yourMenu["eggs"] = 1.75, 1.75
	anotherMenu["bacon"], yourMenu["bacon"] = 3.22, 3.22
	anotherMenu["sausage"], yourMenu["sausage"] = 1.89, 1.89
	fmt.Printf("anotherMenu value:%#v\n", anotherMenu)
	fmt.Printf("anotherMenu value:%#v\n", yourMenu)

}
