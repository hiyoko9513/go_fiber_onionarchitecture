package utils

import (
	"time"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

// Now return current time of UTC
func Now() time.Time {
	return time.Now().UTC()
}

// JstNow return current time of JST
func JstNow() time.Time {
	return time.Now().In(jst)
}

// ToRFC3339 make the time in the form of "RFC3339"
func ToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

// LoadTimezone load local timezone
func LoadTimezone(tz string) {
	time.Local = timezone(tz)
}

// timezone get a time location
func timezone(tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	return loc
}
