package main

import "fmt"

func sum(nums ...int) { // Here’s a function that will take an arbitrary number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Variadic functions can be called with any number of trailing arguments.
	// For example, fmt.Println is a common variadic function.
	// Here’s a function that will take an arbitrary number of ints as arguments.
	sum(1, 2)
	sum(1, 2, 3)
	// Another key aspect of functions in Go is their ability to form closures,
	// which we’ll look at next.
}
