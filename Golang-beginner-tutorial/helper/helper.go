package helper

import "strings"

// capitalize the first letter of the function's name
// to expose it to the package called
func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) bool {
	isValidName := len(firstName)>= 2 && len(lastName) >= 2
	isValidEmail :=	strings.Contains(email,"@")
	isValidTickets :=	userTickets > 0 && userTickets < remainingTickets

	isValid := isValidEmail && isValidName && isValidTickets

	return isValid
}