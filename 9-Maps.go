package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // To create an empty map, use the builtin make: make(map[key-type]val-type).

	m["k1"] = 7 // Set key/value pairs using typical name[key] = val syntax.
	m["k2"] = 14

	fmt.Println("map:", m) // Printing a map with e.g. fmt.Println will show all of its key/value pairs.

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // The builtin len returns the number of key/value pairs when called on a map.

	delete(m, "k2") // The builtin delete removes key/value pairs from a map.
	fmt.Println("map:", m)

	clear(m) // Clear map
	fmt.Println("map:", m)

	key, pres := m["k2"] // The optional second return value when getting a value from a map indicates if the key was present in the map.
	fmt.Println("key:", key, "present:", pres)

	n := map[string]int{"foo": 1, "bar": 2} // Declare and initialize a new map in the same line.
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	}
}
