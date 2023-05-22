package types

import (
	"fmt"
	"testing"
)

func TestDuration(t *testing.T) {
	// string to nanoseconds
	durations := map[string]int64{
		"10s": 10 * 1_000_000_000,
		"10":  10 * 1_000_000_000,
		"1m":  60 * 1_000_000_000,
		"2m":  120 * 1_000_000_000,
		"2h":  60 * 60 * 2 * 1_000_000_000,
		"1d":  60 * 60 * 24 * 1_000_000_000,
		"1w":  60 * 60 * 24 * 7 * 1_000_000_000,
	}

	for durationString, expected := range durations {
		var duration ZBXDuration
		err := duration.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", durationString)))
		if err != nil {
			t.Error(err)
		}

		if expected != int64(duration) {
			t.Errorf("Expected nano seconds %d but got %d for duration %q",
				expected,
				int64(duration),
				durationString)
		}
	}
}
