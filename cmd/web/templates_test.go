package main

import (
	"testing"
	"time"

	"troc.amanya/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 12, 5, 16, 17, 0, 0, time.UTC),
			want: "05 Dec 2024 at 16:17",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 12, 5, 16, 17, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "05 Dec 2024 at 15:17",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}