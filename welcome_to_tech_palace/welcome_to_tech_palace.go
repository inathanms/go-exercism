package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	welcomeMsgWithBorder := border + "\n" + welcomeMsg + "\n" + border
	return welcomeMsgWithBorder
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMsgWithWhiteSpaces := strings.ReplaceAll(oldMsg, "*", "")
	cleanedMsg := strings.TrimSpace(cleanedMsgWithWhiteSpaces)
	return cleanedMsg
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("Welcome!", 10))
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println(CleanupMessage(message))
}
