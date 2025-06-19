package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"

	const conferenceTickets = 50

	var remainingTickets = 50

	fmt.Printf("Hello, World!\n")

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("Total available tickets of %v are %v \n", conferenceTickets, remainingTickets)
}
