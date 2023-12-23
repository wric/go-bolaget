package bolaget

import (
	"go-bolaget/pkg/systembolaget"
	"strconv"
	"strings"
	"time"
)

const NEW_YEARS_EVE = "Nyårsafton"

func ParseSiteName(s *systembolaget.SystembolagetSite) string {
	if s.DisplayName != "" {
		return s.DisplayName
	}
	return s.Alias
}

func MapSite(s *systembolaget.SystembolagetSite) BolagetSite {
	name := ParseSiteName(s)

	address := s.StreetAddress + ", " + s.City
	if s.Postalcode != "" {
		address = s.StreetAddress + ", " + s.Postalcode + " " + s.City
	}

	lat := strconv.FormatFloat(s.Position.Latitude, 'f', -1, 64)
	lng := strconv.FormatFloat(s.Position.Longitude, 'f', -1, 64)

	openingHours := make([]BolagetOpeningHour, len(s.OpeningHours))
	for i, o := range s.OpeningHours {
		openingHours[i] = MapOpeningHour(&o, name)
	}

	return BolagetSite{
		SiteId:       s.SiteId,
		Name:         name,
		Address:      strings.ToLower(address),
		Position:     lat + "," + lng,
		OpeningHours: openingHours,
	}
}

func MapOpeningHour(oh *systembolaget.SystembolagetOpeningHour, store string) BolagetOpeningHour {
	date, err := time.Parse("2006-01-02T00:00:00", oh.Date)
	if err != nil {
		panic(err)
	}

	openTo := oh.OpenTo[:5]

	return BolagetOpeningHour{
		IsClosed:         !isOpen(openTo),
		IsDeviant:        isDeviant(date, openTo),
		ShortDescription: shortDescription(date, openTo),
		FullDescription:  fullDescription(store, date, openTo, oh.Reason),
	}
}

func shortDescription(date time.Time, openTo string) string {
	shortDay := date.Local().Weekday().String()[:3]

	if isOpen(openTo) {
		return shortDay + " " + openTo
	}

	return shortDay
}

func fullDescription(storeName string, date time.Time, openTo string, reason string) string {
	longDay := date.Weekday().String()
	longDate := date.String()[:10]

	if isOpen(openTo) {
		return "Systembolaget " + storeName + " is open to " + openTo + " on " + longDay + " (" + longDate + ")."
	}

	reason = parseReason(date, reason)

	if reason != "" {
		return "Systembolaget " + storeName + " is closed on " + longDay + " (" + longDate + ") due to " + reason + "."
	}

	return "Systembolaget " + storeName + " is closed on " + longDay + " (" + longDate + ")."
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
	return t.Weekday() == 6
}

func isWeekday(t time.Time) bool {
	return t.Weekday() > 0 && t.Weekday() < 6
}

func isNewYearsEve(t time.Time) bool {
	// is utc really needed?
	return t.UTC().Month() == 11 && t.UTC().Day() == 31
}
