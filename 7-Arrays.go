package main

import "fmt"

func main() {
	var a [5]int // Array of 5 integers
	a[4] = 100   // Set the 5th element to 100
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	var twoD [2][3]int // 2D array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
