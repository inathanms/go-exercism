package main

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMsg := fmt.Sprintf("Welcome to my party, %s!", name)
	return welcomeMsg
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	happyBirthdayMsg := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return happyBirthdayMsg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMsg := Welcome(name)
	tableMsg := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	neighborMsg := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	assignTableMsg := strings.Join([]string{welcomeMsg, tableMsg, neighborMsg}, "\n")
	return assignTableMsg
}

func main() {
	fmt.Println(Welcome("Christiane"))
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
