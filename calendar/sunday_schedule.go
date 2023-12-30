package calendar

import (
	"time"

	"github.com/rickar/cal/v2"
)

var (
	months         = []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}
	SundaySchool   []*cal.Holiday
	QuorumMeetings []*cal.Holiday
)

func init() {
	for _, m := range months {
		SundaySchool = append(SundaySchool,
			&cal.Holiday{
				Name:    "Sunday School",
				Type:    cal.ObservanceReligious,
				Month:   m,
				Weekday: time.Sunday,
				Offset:  1,
				Func:    cal.CalcWeekdayOffset,
			},
			&cal.Holiday{
				Name:    "Sunday School",
				Type:    cal.ObservanceReligious,
				Month:   m,
				Weekday: time.Sunday,
				Offset:  3,
				Func:    cal.CalcWeekdayOffset,
			},
		)

		QuorumMeetings = append(QuorumMeetings,
			&cal.Holiday{
				Name:    "Quorum Meetings",
				Type:    cal.ObservanceReligious,
				Month:   m,
				Weekday: time.Sunday,
				Offset:  2,
				Func:    cal.CalcWeekdayOffset,
			},

			&cal.Holiday{
				Name:    "Quorum Meetings",
				Type:    cal.ObservanceReligious,
				Month:   m,
				Weekday: time.Sunday,
				Offset:  2,
				Func:    cal.CalcWeekdayOffset,
			},
		)
	}
}
