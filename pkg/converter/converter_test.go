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
