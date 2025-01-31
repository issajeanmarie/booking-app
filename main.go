package main

import (
	"booking-app/config"
	"booking-app/utils"
	"fmt"
)

func main(){
	welcomeMessage := utils.WelcomeMessage()
	fmt.Println((welcomeMessage))
	var bookings []uint
	remainingTickets := config.SitsAvailable

	for {
		// break when no tickets available
		if(remainingTickets <= 0){
			break
		}

		// break when user enters more tickets than available
		_, _, _, tickets := utils.Input(remainingTickets)
		if(tickets > remainingTickets){
			break
		}

		remainingTickets = remainingTickets - tickets
		bookings = append(bookings, tickets)

	}

	fmt.Println("BOOKINGS: ", bookings)
}