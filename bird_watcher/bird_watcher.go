package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0
	for i := range birdsPerDay {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysAWeek := 7
	if week <= 0 || week > len(birdsPerDay)/daysAWeek {
		return 0
	}
	birdsPerDayInWeek := birdsPerDay[(week-1)*daysAWeek : week*daysAWeek]
	return TotalBirdCount(birdsPerDayInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	fmt.Println(FixBirdCountLog([]int{2, 5, 0, 7, 4, 1}))
}
