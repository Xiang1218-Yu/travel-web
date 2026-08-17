package main

import (
	"fmt"
	"time"
)

type travelDate struct {
	time.Time
}

func parseTravelDate(value string) (travelDate, error) {
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return travelDate{}, fmt.Errorf("invalid travel date %q", value)
	}
	return travelDate{Time: parsed.UTC()}, nil
}

func (d travelDate) after(other travelDate) bool  { return d.Time.After(other.Time) }
func (d travelDate) before(other travelDate) bool { return d.Time.Before(other.Time) }
