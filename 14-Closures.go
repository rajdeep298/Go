package main

func intSeq() func() int { // This function intSeq returns another function, which we define anonymously in the body of intSeq.
	i := 0
	return func() int { // The returned function closes over the variable i to form a closure.
		i++
		return i
	}
}
func main() {
	nextInt := intSeq() // We call intSeq, assigning the result (a function) to nextInt.

	println(nextInt()) // See the effect of the closure by calling nextInt a few times.
	println(nextInt())
	println(nextInt())

	newInts := intSeq() // To confirm that the state is unique to that particular function, create and test a new one.
	println(newInts())
}
