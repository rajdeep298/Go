package main

func plus(a int, b int) int { // Here’s a function that takes two ints and returns their sum as an int.
	return a + b
}

func plusPlus(a, b, c int) int { // When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
	return a + b + c
}

func main() {

	res := plus(1, 2) // Call a function just as you’d expect, with name(args).
	println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	println("1+2+3 =", res)
}
