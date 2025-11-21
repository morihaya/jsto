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
