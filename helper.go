package main

import "strings"

func validateInput(firstName, lastName, email string, userTickets uint, remainingTickets uint, city string) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidCity := city == "Singapore" || city == "London" || city == "singapore" || city == "london"

	return isValidName, isValidEmail, isValidTicketNumber, isValidCity
}
