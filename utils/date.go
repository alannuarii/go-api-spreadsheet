package utils

import (
	"time"
)

func YesterdayDate() string {
	today := time.Now()

	yesterday := today.AddDate(0, 0, -1)

	yesterdayFormatted := yesterday.Format("2006-01-02")

	return yesterdayFormatted
}


func YearMonth() string {
	now := time.Now()

	currentYearMonth := now.Format("2006-01")

	return currentYearMonth
}