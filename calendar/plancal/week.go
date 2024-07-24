package plancal

import (
	"fmt"
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/ag7if"
	"github.com/derhabicht/planning-calendar/calendar/cap"
)

type Week struct {
	calendar calendar.Calendar
	monday   date.Date
	thursday date.Date
}

func NewWeek(cal calendar.Calendar, d date.Date) Week {
	var monday date.Date
	if d.Weekday() == time.Sunday {
		monday = d.Add(-6)
	} else {
		monday = d.Add(-1 * (int(d.Weekday()) - int(time.Monday)))
	}

	thursday := monday.Add(3)

	return Week{
		calendar: cal,
		monday:   monday,
		thursday: thursday,
	}
}

func NewFromISOWeek(cal calendar.Calendar, year, isoweek int) Week {
	// 4 Jan is guaranteed to be a part of week 1, so we use 4 Jan of year as our reference date.
	refDate := date.New(year, time.January, 4)

	// Advance the ref date by isoweek-1 weeks to give us a date within the week that we want.
	d := refDate.Add(7 * (isoweek - 1))

	return NewWeek(cal, d)
}

func (w Week) Trimester() calendar.Trimester {
	fy := cap.ComputeFiscalYear(w.thursday)
	tri := cap.ComputeTrimester(w.thursday)

	return NewTrimester(w.calendar, fy, tri)
}

func (w Week) FiscalQuarter() calendar.Quarter {
	fy := cap.ComputeFiscalYear(w.thursday)
	qtr := cap.ComputeQuarter(w.thursday)

	return NewQuarter(w.calendar, fy, qtr, FiscalQuarter)
}

func (w Week) Quarter() calendar.Quarter {
	year, qtr := ag7if.ComputeQuarter(w.thursday)

	return NewQuarter(w.calendar, year, qtr, CalendarQuarter)
}

func (w Week) Sprint() calendar.Sprint {
	return NewSprint(w.calendar, w.monday)
}

func (w Week) FyWeek() (int, int) {
	fy := cap.ComputeFiscalYear(w.thursday)
	wk1 := calendar.ComputeNearestMonday(date.New(fy-1, time.October, 1))

	return fy, (w.monday.Sub(wk1) / 7) + 1
}

func (w Week) ISOWeek() (int, int, cards.Card) {
	year, week := w.monday.ISOWeek()
	card, err := ag7if.ComputeWeekPlayingCard(week)
	if err != nil {
		panic(errors.WithMessage(err, "unexpected error when computing ISO Week"))
	}

	return year, week, card
}

func (w Week) Next() calendar.Week {
	return NewWeek(w.calendar, w.monday.Add(7))
}

func (w Week) Monday() calendar.Day {
	return NewDay(w.calendar, w.monday)
}

func (w Week) String() string {
	yr, week := w.monday.ISOWeek()
	return fmt.Sprintf("CY%dW%d", yr, week)
}
