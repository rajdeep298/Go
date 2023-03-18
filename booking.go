package main

import "fmt"

func main() {
	var total_ticket int
	fmt.Println("Enter the total number of tickets")
	fmt.Scan(&total_ticket)
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

		//checking the user input coreectness
		if buy > remaining_ticket {
			fmt.Printf("\nTickets left %d.\nYou cannot book %d tickets\n", remaining_ticket, buy)
			continue
		}

		//storing details in an slice
		bookings = append(bookings, f_name+" "+l_name)
		remaining_ticket = remaining_ticket - buy //calculating remaining tickets

		//printing basic details
		fmt.Println(bookings)
		fmt.Println("Remaining tickets:", remaining_ticket)

		//checking if tickets are remaining or not

		//Method-1
		if remaining_ticket == 0 {
			//now end the program
			fmt.Println("No remaining tickets")
			break
		}

		//Method-2
		// var not_ticket bool = remaining_ticket == 0
		// not_ticket := remaining_ticket == 0
		// if not_ticket {
		// 	//now end the program
		// 	fmt.Println("No remaining tickets")
		// 	break
		// }
	}

}
