package utils

import (
	"fmt"
	"time"
)

type DateTimeFormat struct {
	DDMMYYYY          bool
	YYYYMMDD          bool
	AddTime           bool
	AddTimeAndSeconds bool
	Dashed            bool
}

func CalendarAddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func CalendarAddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

func GetTime(format DateTimeFormat, t time.Time) string {
	// format pattern: Mon Jan 2 15:04:05 MST 2006

	tm := ""
	if format.AddTime {
		tm = "15:04"
	}
	if format.AddTimeAndSeconds {
		tm = "15:04:05"
	}

	return t.Format(tm)
}

func GetDateTime(format DateTimeFormat, t time.Time) string {
	// format pattern: Mon Jan 2 15:04:05 MST 2006

	tm := ""
	if format.AddTime {
		tm = " 15:04"
	}
	if format.AddTimeAndSeconds {
		tm = " 15:04:05"
	}

	sep := "/"
	if format.Dashed {
		sep = "-"
	}

	if format.DDMMYYYY {
		s := fmt.Sprintf("02%s01%s2006", sep, sep)
		return t.Format(s + tm)
	}

	if format.YYYYMMDD {
		s := fmt.Sprintf("2006%s01%s02", sep, sep)
		return t.Format(s + tm)
	}

	panic("Invalid format")
}
