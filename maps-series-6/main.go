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

	// Best way to identify a nil map is to use this expression below
	fmt.Println(nilMap == nil)
	isNil := nilMap == nil
	if isNil {
		fmt.Println("Map nilMap is nil")
	}

	//add key and value to a nil map
	nilMap = map[string]bool{
		"God is good": true,
	}

	nilMap["Jesus is alive"] = true
	nilMap["The Holy Spirit is real"] = true
	nilMap["Lucifer is partly good!"] = true

	fmt.Printf("The nilMap now contains: %#v", nilMap)

}
