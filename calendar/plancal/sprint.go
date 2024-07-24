package plancal

import (
	"fmt"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/ag7if"
)

type Sprint struct {
	calendar calendar.Calendar
	year     int
	sprint   ag7if.S
}

func NewSprint(cal calendar.Calendar, d date.Date) Sprint {
	year, sprint := ag7if.ComputeSprint(d)

	return Sprint{
		calendar: cal,
		year:     year,
		sprint:   sprint,
	}
}

func (s Sprint) FirstWeek() (calendar.Week, int) {
	weekNums := ag7if.ComputeWeekNumbers(s.sprint)

	week := NewFromISOWeek(s.calendar, s.year, weekNums[0])

	return week, len(weekNums)
}

func (s Sprint) Short() string {
	return s.sprint.String()
}

func (s Sprint) String() string {
	return fmt.Sprintf("CY%d%s", s.year, s.sprint)
}

func (s Sprint) Full() string {
	return fmt.Sprintf("CY%d, Sprint %s ", s.year, s.sprint)
}
