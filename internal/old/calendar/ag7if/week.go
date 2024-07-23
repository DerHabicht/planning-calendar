package ag7if

import (
	"fmt"
	"strings"
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
	"github.com/nathan-osman/go-sunrise"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/cap"
	"github.com/derhabicht/planning-calendar/calendar/events"
	"github.com/derhabicht/planning-calendar/internal/config"
)

var WeekdayLetters = map[time.Weekday]string{
	time.Monday:    "M",
	time.Tuesday:   "T",
	time.Wednesday: "W",
	time.Thursday:  "H",
	time.Friday:    "F",
	time.Saturday:  "S",
	time.Sunday:    "U",
}

type Week struct {
	monday          date.Date
	fiscalYear      cap.FiscalYear
	fyTrimester     cap.Trimester
	fyQuarter       cap.Quarter
	ag7ifQuarter    Quarter
	ag7ifSprint     Sprint
	isoWeek         int
	card            cards.Card
	currentDay      date.Date
	location        *time.Location
	holidayCalendar *events.HolidayCalendar
	solsticeTable   events.SolsticeTable
}

func NewWeek(d date.Date, loc *time.Location, hc *events.HolidayCalendar, st events.SolsticeTable) *Week {
	_, isoWeek := d.ISOWeek()

	fiscalYear := cap.NewFiscalYearFromDate(d)

	var monday date.Date
	if d.Weekday() == time.Sunday {
		monday = d.Add(-6)
	} else {
		monday = d.Add(-1 * (int(d.Weekday()) - 1))
	}

	thursday := monday.Add(3)

	fyTrimester := cap.NewTrimester(thursday)
	fyQuarter := cap.NewQuarter(thursday)
	ag7ifQuarter := NewQuarter(monday)
	ag7ifSprint := NewSprint(monday)

	return &Week{
		monday:          monday,
		fiscalYear:      fiscalYear,
		fyTrimester:     fyTrimester,
		fyQuarter:       fyQuarter,
		ag7ifQuarter:    ag7ifQuarter,
		ag7ifSprint:     ag7ifSprint,
		isoWeek:         isoWeek,
		card:            GetWeekCard(isoWeek),
		location:        loc,
		holidayCalendar: hc,
		solsticeTable:   st,
	}
}

func (w *Week) CurrentDay() (date.Date, error) {
	if w.currentDay.Sub(w.monday) >= 7 {
		return w.monday, EndOfWeekError{}
	}

	return w.currentDay, nil
}

func (w *Week) CurrentDayStr(month bool) (string, error) {
	if w.currentDay.Sub(w.monday) >= 7 {
		return "", EndOfWeekError{}
	}

	if month || w.currentDay.Day() == 1 {
		return fmt.Sprintf("%s %d",
			strings.ToUpper(w.currentDay.Month().String())[:3],
			w.currentDay.Day(),
		), nil
	}

	return fmt.Sprintf("%d", w.currentDay.Day()), nil
}

func (w *Week) CurrentDaySunriseSunset() (time.Time, time.Time, error) {
	homeLat := config.GetFloat64("home_location.lat")
	homeLong := config.GetFloat64("home_location.long")

	sr, ss := sunrise.SunriseSunset(homeLat, homeLong, w.currentDay.Year(), w.currentDay.Month(), w.currentDay.Day())

	return sr.In(w.location), ss.In(w.location), nil
}

func (w *Week) IsCurrentDaySolstice() (bool, calendar.Solstice) {
	return w.solsticeTable.IsSolstice(w.currentDay)
}

func (w *Week) Next() (date.Date, error) {
	w.currentDay = w.currentDay.Add(1)

	if w.currentDay.Sub(w.monday) >= 7 {
		return w.monday, EndOfWeekError{}
	}

	return w.currentDay, nil
}

func (w *Week) Reset() date.Date {
	w.currentDay = w.monday
	return w.currentDay
}

func (w *Week) FyTrimester() string {
	return w.fyTrimester.String()
}

func (w *Week) FyQuarter() string {
	return w.fyQuarter.String()
}

func (w *Week) FyWeek() string {
	wk, err := w.fiscalYear.FyWeek(w.monday)
	if err != nil {
		panic(fmt.Errorf("CAP Fiscal Year was improperly initialized for week %d: %s", w.isoWeek, err))
	}

	return fmt.Sprintf("W%02d", wk)
}

func (w *Week) CyQuarter() string {
	return w.ag7ifQuarter.String()
}

func (w *Week) Sprint() string {
	return w.ag7ifSprint.String()
}

func (w *Week) IsoWeek() (string, cards.Card) {
	return fmt.Sprintf("W%02d", w.isoWeek), w.card
}

func (w *Week) IsCurrentDayHoliday() (bool, bool, calendar.Holiday) {
	return w.holidayCalendar.IsHoliday(w.currentDay.Local())
}

type EndOfWeekError struct{}

func (eow EndOfWeekError) Error() string {
	return "reached the end of the week"
}
