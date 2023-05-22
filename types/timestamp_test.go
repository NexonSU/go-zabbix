package types

import (
	"testing"
	"time"
)

func TestUnixTimestamp(t *testing.T) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error(err)
	}

	tests := map[string]time.Time{
		"0":          time.Date(1970, 1, 1, 0, 0, 0, 0, loc),
		"1683642493": time.Date(2023, 5, 9, 14, 28, 13, 0, loc),
	}

	for value, expected := range tests {
		var zbxtime ZBXUnixTimestamp
		err := zbxtime.UnmarshalJSON([]byte("\"" + value + "\""))
		if err != nil {
			t.Error(err)
		}

		if time.Time(zbxtime) != expected {
			t.Errorf(
				"Expected '%v' but got '%v' for %q date and time",
				expected,
				time.Time(zbxtime),
				value)
		}
	}
}
