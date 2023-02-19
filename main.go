package main

import (
	"fmt"
	"strings"
)

var eventName = "DevOops event"

const eventTickets = 50

var remainingTickets uint = 50

// empty slice:
var bookings []string

func main() {

	greetUsers(eventName, eventTickets, remainingTickets)

	for {
		// Get user data:
		firstName, lastName, email, userTickets := getUserInput()
		// Input validation:
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			var firstNames = getFirstNames(bookings)
			fmt.Printf("List of current atendees: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// if remainingTickets are 0, finish the loop to exit the program.
				fmt.Printf("Our %v event is booked out\n", eventName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid email address")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you entered is invalid. Currently we have %v tickets left, you can't book %v tickets. Please, try again.\n", remainingTickets, userTickets)
			}
			continue
		}
	}
}

func greetUsers(eventName string, eventTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking application.\n", eventName)
	fmt.Printf("We have a total of %v tickets and %v are still avaliable.\n", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// func + function_name (input_var) output_var {code}
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// "_" is a blanck identifier to ignore a variable you don't want to use.
	for _, booking := range bookings {
		// strings.Fields() splits the string with white space as separator and returns a slice with the split elements.
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets <= remainingTickets && userTickets > 0
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please complete the information requested below.")

	// Ask user for some data:
	// Scan function can assign the user input into the firstName variable because it has a pointer to its memory adress.

	fmt.Println("First name:")
	fmt.Scan(&firstName)

	fmt.Println("Last name:")
	fmt.Scan(&lastName)

	fmt.Println("Email adress:")
	fmt.Scan(&email)

	fmt.Println("How many tikets do you want to get?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v, you have booked %v tickets. You will recieve a confirmation email at %v\n", firstName, userTickets, email)
	fmt.Printf("There are currently %v tickets remaining for the %v\n", remainingTickets, eventName)
}
