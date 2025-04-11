package main

import "fmt"

func main() {
	total_ticket := 500
	remaining_ticket := total_ticket
	var bookings []string
	var f_name string
	var l_name string
	var buy int
	for {
		//taking basic inputs
		fmt.Println("Enter the first name of buyer:")
		fmt.Scan(&f_name)
		fmt.Println("\nEnter the last name of buyer:")
		fmt.Scan(&l_name)
		fmt.Println("\nEnter the amount of tickets buyer wants to book:")
		fmt.Scan(&buy)

		//storing details in an slice
		bookings = append(bookings, f_name+" "+l_name)
		remaining_ticket = remaining_ticket - buy //calculating remaining tickets

		//printing basic details
		fmt.Println(bookings)
		fmt.Println("Remaining tickets:", remaining_ticket)

		//checking if tickets are remaining or not

	}

}
