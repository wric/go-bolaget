package lib

import (
	"go-bolaget/models"
	"time"
)

const NEW_YEARS_EVE = "Nyårsafton"

func MapOpeningHour(store string, oh *models.SystembolagetOpeningHour) models.BolagetOpeningHour {
	openTo := oh.OpenTo[:5]

	layout := "2001-01-01"
	date, err := time.Parse(layout, oh.Date)
	if err != nil {
		panic(err)
	}

	return models.BolagetOpeningHour{}
}

func parseReason(d time.Time, reason string) string {
	if isNewYearsEve(d) {
		return NEW_YEARS_EVE
	}

	if reason == "0" || reason == "-" {
		return ""
	}

	return reason
}

func isDeviant(d time.Time, openTo string) bool {
	if isSunday(d) {
		// Sites are always closed on Sundays.
		return false
	}

	if !isOpen(openTo) {
		// Sites are usually open all days other than Sunday.
		return true
	}

	if isNewYearsEve(d) {
		// New years eve is treated as a Saturday.
		return isWeekday(d)
	}

	return false
}

func isOpen(openTo string) bool {
	return openTo != "00:00"
}

func isSunday(t time.Time) bool {
	return t.Day() == 6
}

func isWeekday(t time.Time) bool {
	return t.Day() > 0 && t.Day() < 6
}

func isNewYearsEve(t time.Time) bool {
	// is utc really needed?
	return t.UTC().Month() == 11 && t.UTC().Day() == 31
}
