package converter

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	fixedTime := time.Date(2023, 10, 27, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		time     time.Time
		location string
		want     string
		wantErr  bool
	}{
		{
			name:     "UTC",
			time:     fixedTime,
			location: "UTC",
			want:     "2023/10/27 10:00:00",
			wantErr:  false,
		},
		{
			name:     "New York (EDT/EST)",
			time:     fixedTime,
			location: "America/New_York",
			want:     "2023/10/27 06:00:00", // UTC-4
			wantErr:  false,
		},
		{
			name:     "Tokyo (JST)",
			time:     fixedTime,
			location: "Asia/Tokyo",
			want:     "2023/10/27 19:00:00", // UTC+9
			wantErr:  false,
		},
		{
			name:     "Invalid Location",
			time:     fixedTime,
			location: "Invalid/Location",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.time, tt.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFromJST(t *testing.T) {
	// Mocking time is tricky here because ConvertFromJST uses time.Now().
	// For simplicity, we will just test that it doesn't return error for valid input
	// and returns error for invalid input.
	// A more robust test would require dependency injection for time.Now().

	tests := []struct {
		name           string
		jstTimeStr     string
		targetLocation string
		wantErr        bool
	}{
		{
			name:           "Valid Time UTC",
			jstTimeStr:     "13:00",
			targetLocation: "UTC",
			wantErr:        false,
		},
		{
			name:           "Valid Time EDT",
			jstTimeStr:     "13:00",
			targetLocation: "America/New_York",
			wantErr:        false,
		},
		{
			name:           "Invalid Time Format",
			jstTimeStr:     "25:00",
			targetLocation: "UTC",
			wantErr:        true,
		},
		{
			name:           "Invalid Location",
			jstTimeStr:     "13:00",
			targetLocation: "Invalid/Location",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertFromJST(tt.jstTimeStr, tt.targetLocation)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFromJST() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
