package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets, city := userInput()
		isValidName, isValidEmail, isValidTicketNumber, isValidCity := validateInput(firstName, lastName, email, userTickets, remainingTickets, city)

		if isValidEmail && isValidName && isValidTicketNumber && isValidCity {
			bookTicket(firstName, lastName, email, userTickets)
			firstNames := getNames()
			fmt.Printf("booking details: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out for", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name must be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid. Please enter a valid email address.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets must be greater than 0 and less than or equal to remaining tickets.")
			}
			if !isValidCity {
				fmt.Println("City must be either Singapore or London.")
			}
		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to the %v booking application! available tickets are %v\n", conferenceName, remainingTickets)
}

func getNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func userInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var city string

	fmt.Println("please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("please enter number of tickets you want to book:")
	fmt.Scan(&userTickets)
	fmt.Println("please enter city")
	fmt.Scan(&city)

	return firstName, lastName, email, userTickets, city
}

func bookTicket(firstName, lastName, email string, userTickets uint) {
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets for", conferenceName)
	fmt.Printf("You will receive a confirmation email on %v at the end of the booking process.\n", email)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)
}
