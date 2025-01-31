package utils

import "fmt"

func Input(remainingTickets uint)(string, string, string, uint) {
	var fName string
	var lName string
	var email string
	var tickets uint

	fmt.Print("Enter first name: ")
	fmt.Scan(&fName)
	
	fmt.Print("Enter last name: ")
	fmt.Scan(&lName)

	fmt.Print(("Enter email: "))
	fmt.Scan(&email)
	
	fmt.Printf("Enter tickets (%v): ", remainingTickets)
	fmt.Scan(&tickets)

	return fName, lName, email, tickets
}