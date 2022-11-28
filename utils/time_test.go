package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {

	type Test struct {
		Time Time
	}

	dates := []string{
		`2021-09-09 07:52:34 UTC`,
		`2020-09-23 21:20:05 UTC`,
		`2020-09-23 21:20:04 UTC`,
	}

	for _, date := range dates {
		a := Test{}
		data := fmt.Sprintf("{\"Time\": \"%s\"}", date)
		t.Logf("testing: %q", data)
		err := json.Unmarshal([]byte(data), &a)
		if err != nil {
			t.Fatal(err)
		}
		if a.Time.String() != date {
			t.Fatalf("error: expected %q, got %q", date, a.Time.String())
		}
	}
}
