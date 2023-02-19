package main

import (
	"fmt"
	"strings"
)

func main() {
	eventName := "DevOops event"
	const eventTickets = 50
	var remainingTickets uint = 50
	// empty slice:
	var bookings []string

	fmt.Printf("Welcome to our %v booking application.\n", eventName)
	fmt.Printf("We have a total of %v tickets and %v are still avaliable.\n", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please complete the information requested below.")

		// Ask user for their first name
		fmt.Println("First name:")
		// Scan function can now assign the user input into the firstName variable because it has a pointer to its memory adress
		fmt.Scan(&firstName)

		// Ask user for their last name
		fmt.Println("Last name:")
		fmt.Scan(&lastName)

		// Ask user for their email
		fmt.Println("Email adress:")
		fmt.Scan(&email)

		// Ask user how many tikets he wants:
		fmt.Println("How many tikets do you want to get?")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("Array type: %T\n", bookings)
			//fmt.Printf("Array lenght: %v\n", len(bookings))

			fmt.Printf("Thank you %v, you have booked %v tickets. You will recieve a confirmation email at %v\n", firstName, userTickets, email)
			fmt.Printf("There are currently %v tickets remaining for the %v\n", remainingTickets, eventName)

			firstNames := []string{}
			// "_" is a blanck identifier to ignore a variable you don't want to use.
			for _, booking := range bookings {
				// strings.Fields() splits the string with white space as separator and returns a slice with the split elements.
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("List of current atendees: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// if remainingTickets are 0, finish the loop to exit the program.
				fmt.Printf("Our %v event is booked out\n", eventName)
				break
			}
		} else {
			fmt.Printf("We only have %v tickets left, you can't book %v tickets. Please, try again.\n", remainingTickets, userTickets)
			// "continue" causes loop to skip the remainder of its body and retest its condition.
			continue
		}

	}
}
