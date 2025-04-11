package main

import (
	"fmt"
	"strings"
)

// var total_ticket int
// var remaining_ticket = total_ticket// in global variables := will not work

func main() {
	var total_ticket int
	fmt.Println("Enter the total number of tickets")
	fmt.Scan(&total_ticket)
	remaining_ticket := total_ticket
	var bookings []string
	var mails []string
	var f_name string
	var l_name string
	var city string
	var buy int
	var email string
	conf_name := "Go Conference"
	for remaining_ticket > 0 /*or len(bookings) <= 500*/ {
		greetusers(conf_name, remaining_ticket) //incase of global variables we dont need any function parameters

		//taking basic inputs
		fmt.Println("Enter the first name of buyer:")
		fmt.Scan(&f_name)
		fmt.Println("\nEnter the last name of buyer:")
		fmt.Scan(&l_name)
		fmt.Println("\nEnter the amount of tickets buyer wants to book:")
		fmt.Scan(&buy)
		fmt.Println("\nEnter the email of the buyer:")
		fmt.Scan(&email)
		fmt.Println("\nEnter the name of the city")
		fmt.Scan(&city)
		//checking the user input coreectness
		if buy > remaining_ticket {
			fmt.Printf("\nTickets left %d.\nYou cannot book %d tickets\n", remaining_ticket, buy)
			fmt.Printf("\nEnter the correct details.\n")
			continue
		}
		switch city {
		case "Kolkata":
			fmt.Println("Your venue is Kolkata")
		case "Delhi":
			fmt.Println("Your venue is DELHI")
		case "Bengaluru":
			fmt.Println("Your venue is BENGALURU")
		default:
			fmt.Println("Your city is not available")
			continue
		}
		vaild := strings.Contains(email, "@")
		if len(l_name) >= 2 && len(f_name) >= 2 && vaild /*or strings.Contains(email, "@")*/ {

			//storing details in an slice
			bookings = append(bookings, f_name+" "+l_name)
			mails = append(mails, email)
			remaining_ticket = remaining_ticket - buy //calculating remaining tickets

			//printing basic details
			fmt.Println(bookings)
			fmt.Println(mails)
			fmt.Println("Remaining tickets:", remaining_ticket)
		} else {
			println("Enter the details correctly\n")
			continue
		}
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

func greetusers(name string, rem_ticket int) {
	fmt.Printf("\n\nWelcome to the %v\n", name)
	fmt.Printf("%v tickets left. Hurry Up!\n", rem_ticket)
}
