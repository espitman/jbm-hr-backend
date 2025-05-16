package utils

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

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

// GenerateDatesForJalaliMonth returns a slice of time.Time for all dates in the current Jalali month for the last n years
func GenerateDatesForJalaliMonth(years int) []time.Time {
	// Get current Jalali date
	now := ptime.Now()
	currentJalaliMonth := now.Month()
	currentJalaliYear := now.Year()

	// Get first and last day of current Jalali month
	firstDayOfMonth := ptime.Date(currentJalaliYear, currentJalaliMonth, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	// Convert to Gregorian
	firstDayGregorian := firstDayOfMonth.Time()
	lastDayGregorian := lastDayOfMonth.Time()

	var dates []time.Time
	currentYear := time.Now().Year()

	// For each year in the range
	for y := currentYear - years; y <= currentYear; y++ {
		// Create date range for this year
		startDate := time.Date(y, firstDayGregorian.Month(), firstDayGregorian.Day(), 0, 0, 0, 0, time.UTC)
		endDate := time.Date(y, lastDayGregorian.Month(), lastDayGregorian.Day(), 0, 0, 0, 0, time.UTC)

		// Add all dates in the range
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			dates = append(dates, d)
		}
	}

	return dates
}
