package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		//var conf_name = "Go Conference"
		conf_name := "Go Conference" // instead of using var we can use this
		const ticket_price = 50      // const value cant be changed elsewhere in the prog
		// if it was a var then it can be changed elsewhere in the prog
		ticket_num := 100
		fmt.Printf("Conference name:%T,Ticket price:%T\n", conf_name, ticket_price) // %T is used to print the data type name
		fmt.Println("Welcome to", conf_name, "booking application")
		fmt.Println("Get your tickects here to attend")

		// fmt.Println(conf_name)

		fmt.Println("Ticket price:", ticket_price)

		fmt.Printf("Name of the conference:%v\n", conf_name) //%v is used for the value in the variable irrespective of the datatype

		var user string
		// ask user for intput
		// user = "Rajdeep"

		fmt.Scan(&user)
		var num int
		num = 20

		fmt.Println(user)
		fmt.Println(num)

		//fmt.Println(&user)

		remaining_ticket := ticket_num - num //basic arithmatic operation
		fmt.Println(remaining_ticket)

		// var bookings = [50]string{"raj1", "raj2", "raj3", "raj4", "raj5", "raj6", "raj7"}

		var booking []string // declartion of an array
		//booking[0] = "Rajdeep"
		//booking[10] = "Kundu"

		var f_name string
		var l_name string
		fmt.Scan(&f_name)
		fmt.Scan(&l_name)
		//booking[5] = f_name + " " + l_name // fittng sentence in a place
		booking = append(booking, f_name+" "+l_name)
		fmt.Println("The Whole array is:")
		fmt.Printf("%v ", booking)
		fmt.Printf("\n")
		//fmt.Printf("An element:%v\n", booking[5])
		fmt.Printf("type of array:%T", booking)
		fmt.Printf("\n\n")

		firstNames := []string{} // that's how you declare a slice
		for _, book := range booking {
			var names = strings.Fields(book)
			firstNames = append(firstNames, names[0])
		}

	}
}
