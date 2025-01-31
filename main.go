package main

import (
	"booking-app/config"
	"booking-app/utils"
	"fmt"
	"time"
)

var bookings []uint
func main(){
	welcomeMessage := utils.WelcomeMessage()
	fmt.Println((welcomeMessage))
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
		bookTicket(tickets)
		go sendTicket()

	}

}

func bookTicket (tickets uint) {
	bookings = append(bookings, tickets)
}

func sendTicket() {
	time.Sleep(10 * time.Second)
	var emailMessage = "Thank you for booking tickets with us.."
	fmt.Println((emailMessage))
}