package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string                                                // A slice is a dynamically-sized, flexible view into the elements of an array.
	fmt.Println("uninitialized slice:", s, s == nil, len(s) == 0) // The zero value of a slice is nil.

	s = make([]string, 3)
	fmt.Println("Empty slice:", s, "len:", len(s), "cap:", cap(s)) // We can set a value at an index just like with an array.

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	s = append(s, "d") // Append returns a slice containing one or more new values.
	s = append(s, "e", "f")
	fmt.Println("Appended:", s)

	c := make([]string, len(s))
	copy(c, s) // We can also copy one slice to another.
	fmt.Println("Copied:", c)

	l := s[2:5] // Slice operator s[2:5] selects elements 2 through 4 (excluding 5).
	fmt.Println("Slice 1:", l)

	l = s[:5] // Slice up to (but excluding) s[5].
	fmt.Println("Slice 2:", l)

	l = s[2:] // Slice up from (and including) s[2].
	fmt.Println("Slice 3:", l)

	t := []string{"g", "h", "i"} // Declare and initialize a variable for slice in a single line.
	fmt.Println("Declared Slice:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ { // To create a 2D slice, we can declare a slice of slices.
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Slice:", twoD)
}
