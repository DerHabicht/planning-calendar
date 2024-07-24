package calendar

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/derhabicht/planning-calendar/internal/config"
)

func WeekdayLetter(wd time.Weekday) string {
	switch wd {
	case time.Monday:
		return "M"
	case time.Tuesday:
		return "T"
	case time.Wednesday:
		return "W"
	case time.Thursday:
		return "H"
	case time.Friday:
		return "F"
	case time.Saturday:
		return "S"
	case time.Sunday:
		return "U"
	default:
		panic(errors.Errorf("invalid weekday value: %d", wd))
	}
}

func ComputeNearestMonday(d date.Date) date.Date {
	var dd int
	wd := d.Weekday()
	switch {
	case time.Tuesday <= wd && wd <= time.Thursday:
		dd = -1 * (int(wd) - 1)
	case wd >= time.Friday:
		dd = 8 - int(wd)
	case wd == time.Sunday:
		dd = 1
	default:
		dd = 0
	}

	return d.Add(dd)
}

func ComputeNearestThursday(d date.Date) date.Date {
	monday := ComputeNearestMonday(d)
	return monday.Add(3)
}

func ComputeLastDayOfMonth(d date.Date) int {
	switch d.Month() {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if date.New(d.Year(), time.December, 31).YearDay() == 366 {
			return 29
		}

		return 28
	default:
		panic(errors.Errorf("not a valid time.Month value: %d", d.Month()))
	}
}

func GetLocation() *time.Location {
	loc, err := time.LoadLocation(config.GetString("home_location.tz"))
	if err != nil {
		log.Warn().Str("tz", config.GetString("home_location.tz")).Msg("Unable to load time zone from config, defaulting to local system time")
		loc = time.Local
	}

	return loc
}

func TimeToLocalDate(t time.Time) date.Date {
	loc := GetLocation()
	local := t.In(loc)
	return date.New(local.Year(), local.Month(), local.Day())
}

func DateToLocalTime(d date.Date) time.Time {
	loc := GetLocation()
	return d.In(loc)
}
