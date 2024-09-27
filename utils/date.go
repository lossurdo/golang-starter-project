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
	DashSepareted     bool
}

func GetTime(format DateTimeFormat) string {
	t := time.Now() // format pattern: Mon Jan 2 15:04:05 MST 2006

	tm := ""
	if format.AddTime {
		tm = "15:04"
	}
	if format.AddTimeAndSeconds {
		tm = "15:04:05"
	}

	return t.Format(tm)
}

func GetDateTime(format DateTimeFormat) string {
	t := time.Now() // format pattern: Mon Jan 2 15:04:05 MST 2006

	tm := ""
	if format.AddTime {
		tm = " 15:04"
	}
	if format.AddTimeAndSeconds {
		tm = " 15:04:05"
	}

	sep := "/"
	if format.DashSepareted {
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
