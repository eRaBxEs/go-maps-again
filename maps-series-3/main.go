package main

import (
	"fmt"
)

func main() {

	// using string literal to create a map
	inStock := map[string]bool{
		"eggs":    true,
		"bacon":   true,
		"sausage": true,
	}

	dish := "sausage"

	// checking for existence of a given item in the boolean map

	// lame way of doing it
	if _, ok := inStock[dish]; ok {
		fmt.Printf("Yes, %s is in stock", dish)
	} else {
		fmt.Printf("Sorry, we're all out of %s", dish)
	}

	// clean way of doing it
	if inStock[dish] {
		fmt.Printf("Yes, %s is in stock", dish)
	} else {
		fmt.Printf("Sorry, we're all out of %s", dish)
	}

}
