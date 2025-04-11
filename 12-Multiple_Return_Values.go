package main

func nos() (int, int) {
	return 1, 2
}
func main() {
	// Functions can return any number of results.
	// The swap function returns two strings.
	a, b := nos()
	println(a, b)
	_, c := nos()
	println(c)
}
