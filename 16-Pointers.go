package main

import "fmt"

func double(n int) int {
	return 2 * n
}
func pointer(nptr *int) {
	*nptr = 5
}
func main() {
	i := 1

	fmt.Println("Initials:", i)

	fmt.Println("Call by value:", double(i))

	pointer(&i)

	fmt.Println("After changing by pointer:", i)
}
