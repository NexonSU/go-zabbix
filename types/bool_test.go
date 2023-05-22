package types

import "testing"

func TestBool(t *testing.T) {
	tests := map[string]bool{
		"1":     true,
		"true":  true,
		"0":     false,
		"false": false,
	}

	for input, expected := range tests {
		var data ZBXBoolean
		data.UnmarshalJSON([]byte(input))

		if data != ZBXBoolean(expected) {
			t.Errorf("Expected %q to be %t", input, expected)
		}
	}
}
