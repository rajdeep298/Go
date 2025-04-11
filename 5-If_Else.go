package main

import "fmt"

func main() {
	if num := 158; num < 10 {
		fmt.Println("<10=>", num)
	} else if num < 100 {
		fmt.Println(">10=>", num)
	} else {
		fmt.Println("mismatch")
	}
}
