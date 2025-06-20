package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"

	const conferenceTickets = 50

	var remainingTickets = 50
	var userName string
	var userTickets int

	userName = "John Doe"
	userTickets = 2

	fmt.Printf("Hello, World!\n")
	fmt.Printf("Hello, %v\n", userName)
	fmt.Printf("You have, %v tickets\n", userTickets)

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("Total available tickets of %v are %v \n", conferenceTickets, remainingTickets)
}
