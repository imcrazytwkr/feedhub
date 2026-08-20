package models

import (
	"encoding/json"
	"time"
)

type Date time.Time

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time().Format(time.DateOnly))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	timestamp, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}

	*d = Date(timestamp)
	return nil
}
