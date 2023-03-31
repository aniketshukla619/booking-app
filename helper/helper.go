package helper

import "strings"

// function to validate the user input
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {// (bool, bool, bool) is the return type of the return value.
												 
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// returning multiple values in the return statement
	return isValidName , isValidEmail , isValidTicketNumber
}