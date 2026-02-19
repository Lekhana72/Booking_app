package main

import "strings"

func ValidateUserInput(firstname string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// not only can export packages but can also variables and constants
// Three levels of Scope:
// 1. Local Scope: Declaration within  function and within block (e.g for, if-else )
// 2. Package Scope: Declaration outside of function but within package
// 3. Global Scope: Declaration outside all functions and Uppercase first letter
