package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	scheduleTime, _ := time.Parse(layout, date)

	return scheduleTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// July 25, 2019 13:45:00"
	layout := "January 2, 2006 15:04:05"
	dateTime, _ := time.Parse(layout, date)
	hasPassed := dateTime.Compare(time.Now()) == -1
	return hasPassed
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// "Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 2, 2006 15:04:05"
	appointmentTime, _ := time.Parse(layout, date)
	appointmentHour := appointmentTime.Hour()
	isAfternoonAppointment := appointmentHour >= 12 && appointmentHour < 18
	return isAfternoonAppointment
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}

func main() {
	// fmt.Println(Schedule("7/25/2019 13:45:00"))
	// fmt.Println(HasPassed("December 9, 2112 11:59:59"))
	fmt.Println(IsAfternoonAppointment("Friday, March 8, 1974 12:02:02"))
	// fmt.Println(Description("7/25/2019 13:45:00"))
	// fmt.Println(AnniversaryDate())
}
