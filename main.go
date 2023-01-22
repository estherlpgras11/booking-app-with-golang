package main

import "fmt"

func main() {
	eventName := "DevOops event"
	const eventTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application.\n", eventName)
	fmt.Printf("We have a total of %v tickets and %v are still avaliable.\n", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v, you have booked %v tickets. You will recieve a confirmation email at %v\n", firstName, userTickets, email)
	fmt.Printf("There are currently %v tickets remaining for the %v\n", remainingTickets, eventName)
}
