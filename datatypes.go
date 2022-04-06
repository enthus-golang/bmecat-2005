package bmecat

import (
	"time"
)

type Datetime struct {
	Date string `xml:"DATE"`
	Time string `xml:"TIME,omitempty"`
}

func NewDatetime(typ string, t time.Time, withTime bool) *Datetime {
	d := &Datetime{
		Date: t.Format("2006-01-02"),
	}

	if withTime {
		d.Time = t.Format("15:04:05")
	}

	return d
}
