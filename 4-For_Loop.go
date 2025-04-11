package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for j := 0; j < 10; j++ {

		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
