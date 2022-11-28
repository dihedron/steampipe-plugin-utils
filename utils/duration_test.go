package utils

import (
	"encoding/json"
	"testing"
)

func TestSatelliteDuration(t *testing.T) {

	type Test struct {
		Elapsed Duration `json:"elapsed"`
	}

	var tests = []string{
		`{"elapsed": "10s"}`,
		`{"elapsed": "30m"}`,
		`{"elapsed": "30ms"}`,
		`{"elapsed": "10ns"}`,
	}

	for _, test := range tests {
		a := Test{}
		err := json.Unmarshal([]byte(test), &a)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("after unmarshalling: %q (%d)", a.Elapsed.String(), a.Elapsed)
	}
}
