package utils

import "time"

// GenerateDatesForTodayMonthDay returns a slice of time.Time for today's month and day for the last n years
func GenerateDatesForTodayMonthDay(years int) []time.Time {
	today := time.Now()
	todayMonth := today.Month()
	todayDay := today.Day()
	currentYear := today.Year()

	var dates []time.Time
	for y := currentYear - years; y <= currentYear; y++ {
		dates = append(dates, time.Date(y, todayMonth, todayDay, 0, 0, 0, 0, time.UTC))
	}
	return dates
}
