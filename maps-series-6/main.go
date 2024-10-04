package main

import (
	"fmt"
)

func main() {
	// To check if any map is empty; get the length and if it is zero, map is empty
	m := map[string]bool{
		"Go is awesome":             true,
		"I drink and I know things": true,
		"PI IS EXACTLY THREE!":      false,
	}
	fmt.Println(len(m))

	emptyMapOne := map[string]bool{}
	emptyMapTwo := make(map[string]bool)

	fmt.Println("emptyMapOne has a length of", len(emptyMapOne))
	fmt.Println("emptyMapTwo has a length of", len(emptyMapTwo))

	// Creating a nil map
	var nilMap map[string]bool // Note that a nil map will also have a length of zero
	fmt.Println("nilMap has a length of", len(nilMap))

}
