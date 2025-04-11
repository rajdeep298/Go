package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//fmt.Println(time.Saturday)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // Multiple expressions in the same case statement
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch { // Switch without an expression is an alternate way to express if/else logic.
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}) { // Type switch compares types instead of values.
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case float64:
			fmt.Println("I'm a float64")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(1.0)
	whatAmI(1.0 + 1.0i)
}
