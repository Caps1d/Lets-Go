package main

import (
	"github.com/Caps1d/Lets-Go/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	// hd := humanDate(tm)
	//
	// if hd != "17 Mar 2022 at 10:15" {
	// 	t.Errorf("got %q; want %q", hd, "17 Mar 2022 at 10:15")
	// }

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
