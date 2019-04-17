package bmecat

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

const datetimeFormat = "2006-01-02T15:04:05Z07:00"

var datetimeRegExp = regexp.MustCompile(`^(\d{4})(?:\-(0[1-9]|1[0-2])(?:\-(0[1-9]|1[0-9]|2[0-9]|3[0-1])(?:T(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]):(?:([0-5][0-9])(?:\.([0-9]{1,})){0,1}){0,1}(?:([+\-])(?:([0-1][0-9]|2[0-3]):([0-5][0-9]))|Z){0,1}){0,1}){0,1}){0,1}$`)

type Datetime struct {
	time.Time
}

func (d Datetime) MarshalText() ([]byte, error) {
	if y := d.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(datetimeFormat))
	return d.Time.AppendFormat(b, datetimeFormat), nil
}

func (d *Datetime) UnmarshalText(data []byte) error {
	matches := datetimeRegExp.FindStringSubmatch(string(data))
	if matches == nil {
		return nil
	}

	var (
		err error

		year   int
		month  = 1
		day    = 1
		hour   int
		minute int
		second int
		nsec   int
		loc    = time.UTC
	)

	if matches[1] != "" {
		year, err = strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
	}
	if matches[2] != "" {
		month, err = strconv.Atoi(matches[2])
		if err != nil {
			return err
		}
	}
	if matches[3] != "" {
		day, err = strconv.Atoi(matches[3])
		if err != nil {
			return err
		}
	}
	if matches[4] != "" {
		hour, err = strconv.Atoi(matches[4])
		if err != nil {
			return err
		}
	}
	if matches[5] != "" {
		minute, err = strconv.Atoi(matches[5])
		if err != nil {
			return err
		}
	}
	if matches[6] != "" {
		second, err = strconv.Atoi(matches[6])
		if err != nil {
			return err
		}
	}
	if matches[7] != "" {
		nsec, err = strconv.Atoi(matches[7])
		if err != nil {
			return err
		}
	}
	if matches[8] != "" && matches[9] != "" {
		zone, err := strconv.Atoi(matches[9])
		if err != nil {
			return err
		}
		if matches[8] == "-" {
			zone *= -1
		}
		loc = time.FixedZone("", zone*3600)
	}

	d.Time = time.Date(year, time.Month(month), day, hour, minute, second, nsec, loc)
	return nil
}
