package apptime

import (
	"log"
	"strconv"
	"time"
)

const TZ_OFFSET = "+0530"

var ALLOWED_OFFSETS = map[string]string{
	"+0530": "IST",
	"+0000": "UTC",
}

const TIME_FORMAT = "2006-01-02 15:04:05 -0700" // system expects the exact timing for the formatting

// InitTimezone is called during app startup to validate the ENV
func getTimezone() *time.Location {
	tzOffset := TZ_OFFSET

	// 1. Check if mandatory field exists
	if tzOffset == "" {
		tzOffset = "+0000"
		log.Fatal("ERROR: TIMEZONE_OFFSET is mandatory in .env (e.g., +0530)")
	}

	tzName, exists := ALLOWED_OFFSETS[tzOffset]
	if !exists {
		// 3. Throw error if the offset is not in our approved list
		log.Fatalf("ERROR: Invalid TIMEZONE_OFFSET '%s'. Does not match any known zones.", tzOffset)
	}

	// 4. Calculate the fixed zone once
	hours, _ := strconv.Atoi(tzOffset[1:3])
	mins, _ := strconv.Atoi(tzOffset[3:5])
	seconds := (hours*3600 + mins*60)

	if tzOffset[0] == '-' {
		seconds = -seconds
	}

	return time.FixedZone(tzName, seconds)
}

func Time(t time.Time) time.Time {
	fixedZone := getTimezone()
	if fixedZone == nil {
		return t.UTC()
	}
	return t.In(fixedZone)
}

func FormattedTime(t time.Time) string {
	return Time(t).Format(TIME_FORMAT)
}

func CurrentFormattedTime() string {
	return Time(time.Now()).Format(TIME_FORMAT)
}
