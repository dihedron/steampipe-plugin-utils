package utils

import (
	"encoding/json"
	"errors"
	"time"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		temp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(temp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d *Duration) String() string {
	return time.Duration(*d).String()
}
