package bmecat

import (
	"time"
)

type Datetime struct {
	Date string
}

func NewDatetime(t time.Time) *Datetime {
	d := &Datetime{
		Date: t.Format("2006-01-02"),
	}

	return d
}
