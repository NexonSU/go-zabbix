package types

import (
	"encoding/json"
	"testing"
)

func TestVersion(t *testing.T) {
	v600, _ := NewZBXVersion("6.0.0")
	v610, _ := NewZBXVersion("6.1.0")
	v615, _ := NewZBXVersion("6.1.5")

	tests := map[string]*ZBXVersion{
		"6.0.0": v600,
		"6.1.0": v610,
		"6.1.5": v615,
	}

	for input, expected := range tests {
		var version ZBXVersion
		err := json.Unmarshal([]byte("\""+input+"\""), &version)
		if err != nil {
			t.Error(err)
		}

		if expected.String() != version.String() {
			t.Errorf("Expected version %q to be %q", expected.String(), input)
		}
	}
}

func TestVersionCompare(t *testing.T) {
	v500, _ := NewZBXVersion("5.0.0")
	v600, _ := NewZBXVersion("6.0.0")
	v610, _ := NewZBXVersion("6.1.0")
	v615, _ := NewZBXVersion("6.1.5")

	tests := map[string]map[*ZBXVersion]int{
		"5.0.0": {
			v500: 0,
			v600: 1,
			v610: 1,
			v615: 1,
		},
		"6.1.0": {
			v500: -1,
			v600: -1,
			v610: 0,
			v615: 1,
		},
		"6.1.5": {
			v500: -1,
			v600: -1,
			v610: -1,
			v615: 0,
		},
	}

	for version, versions := range tests {
		zbxVersion, err := NewZBXVersion(version)
		if err != nil {
			t.Error(err)
		}

		for compare, expected := range versions {
			result := compare.Compare(zbxVersion)
			if result != expected {
				t.Errorf(
					"Version %q when compared to %q should return %d, but it is %d ",
					zbxVersion,
					compare,
					expected,
					result)
			}
		}
	}
}
