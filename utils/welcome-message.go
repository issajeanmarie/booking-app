package utils

import "booking-app/config"

func WelcomeMessage() string {
	message := "Hi, welcome to " + config.ConferenceName + ". Start booking tickets now."
	return message
}