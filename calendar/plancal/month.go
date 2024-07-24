package plancal

import (
	"fmt"
	"strings"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

type Month struct {
	calendar     calendar.Calendar
	calendarYear int
	month        time.Month
}

func NewMonth(cal calendar.Calendar, y int, m time.Month) Month {
	return Month{
		calendar:     cal,
		calendarYear: y,
		month:        m,
	}
}

func (m Month) Year() int {
	return m.calendarYear
}

func (m Month) Month() time.Month {
	return m.month
}

func (m Month) FirstWeek() calendar.Week {
	week := NewWeek(m.calendar, date.New(m.calendarYear, m.month, 1))

	return &week
}

func (m Month) Short() string {
	return strings.ToUpper(m.month.String()[:3])
}

func (m Month) String() string {
	return fmt.Sprintf("%s %d", m.month, m.calendarYear)
}

func (m Month) Next() calendar.Month {
	if m.month == time.December {
		return NewMonth(m.calendar, m.calendarYear+1, time.January)
	}

	return NewMonth(m.calendar, m.calendarYear, m.month+1)
}
