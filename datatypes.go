package bmecat

import (
	"time"
)

type Date string

func NewDate(t time.Time) *Date {
	d := Date(t.Format("2006-01-02"))

	return &d
}
