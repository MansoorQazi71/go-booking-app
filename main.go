package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	var firstName string
	var lastName string
	var userTickets uint
	var email string
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to ", conferenceName, "booking application")
	for remainingTickets > 0 && len(bookings) < 50 {
		fmt.Println("please enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)
		fmt.Println("please enter number of tickets you want to book:")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {
			bookings = append(bookings, firstName+" "+lastName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("booking details: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out for", conferenceName)
				break
			}
			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets for", conferenceName)
			fmt.Printf("You will receive a confirmation email on %v at the end of the booking process.\n", email)
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)
		} else {
			fmt.Printf("Sorry, we only have %v tickets remaining for %v\n", remainingTickets, conferenceName)
		}
	}

}
