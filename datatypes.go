package bmecat

import (
	"time"
)

type Date struct {
	Date string
}

func NewDate(t time.Time) *Date {
	d := &Date{
		Date: t.Format("2006-01-02"),
	}

	return d
}
