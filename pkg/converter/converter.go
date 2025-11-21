package converter

import (
	"fmt"
	"time"
)

// Convert returns the formatted time string in the specified location.
// It takes a time.Time object and a location string (e.g., "America/New_York", "UTC").
func Convert(t time.Time, location string) (string, error) {
	if location == "UTC" {
		return t.UTC().Format("2006/01/02 15:04:05"), nil
	}

	loc, err := time.LoadLocation(location)
	if err != nil {
		return "", fmt.Errorf("failed to load location %s: %w", location, err)
	}

	return t.In(loc).Format("2006/01/02 15:04:05"), nil
}

// ConvertFromJST parses a JST time string (HH:mm) and converts it to the target location.
// It assumes the current date.
func ConvertFromJST(jstTimeStr string, targetLocation string) (string, error) {
	// Parse the input time string
	layout := "15:04"
	parsedTime, err := time.Parse(layout, jstTimeStr)
	if err != nil {
		return "", fmt.Errorf("invalid time format: %w", err)
	}

	// Get current date in JST
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", fmt.Errorf("failed to load JST location: %w", err)
	}
	now := time.Now().In(jst)

	// Combine current date with parsed time
	t := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, jst)

	return Convert(t, targetLocation)
}
