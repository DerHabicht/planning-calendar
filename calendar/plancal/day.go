package plancal

import (
	"time"

	moon "github.com/IvanMenshykov/MoonPhase"
	"github.com/fxtlabs/date"
	"github.com/nathan-osman/go-sunrise"
	"github.com/pkg/errors"
	"github.com/rickar/cal/v2"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/internal/config"
)

type Day struct {
	calendar calendar.Calendar
	date     date.Date
	sunrise  time.Time
	sunset   time.Time
	phase    calendar.MoonPhase
}

func NewDay(cal calendar.Calendar, d date.Date) Day {
	lat := config.GetFloat64("home_location.lat")
	long := config.GetFloat64("home_location.long")
	loc := calendar.GetLocation()
	sr, ss := sunrise.SunriseSunset(lat, long, d.Year(), d.Month(), d.Day())

	m := moon.New(d.In(loc))
	phase, err := calendar.ParseMoonPhase(m.PhaseName())
	if err != nil {
		panic(errors.WithMessage(err, "unexpected error encountered while computing the phase of the moon"))
	}

	return Day{
		calendar: cal,
		date:     d,
		sunrise:  sr.In(loc),
		sunset:   ss.In(loc),
		phase:    phase,
	}
}

func (d Day) Date() date.Date {
	return d.date
}

func (d Day) ISODate() string {
	return d.date.FormatISO(4)
}

func (d Day) IsHoliday() (bool, bool, calendar.Holiday) {
	return d.calendar.HolidayCalendar().IsHoliday(d.date)
}

func (d Day) IsSolstice() calendar.Solstice {
	return d.calendar.SolsticeTable().IsSolstice(d.date)
}

func (d Day) MoonPhase() calendar.MoonPhase {
	return d.phase
}

func (d Day) OrdinalDay() int {
	return d.date.YearDay()
}

func (d Day) MJD() int {
	return int(cal.ModifiedJulianDate(calendar.DateToLocalTime(d.date)))
}

func (d Day) Sunrise() time.Time {
	return d.sunrise
}

func (d Day) Sunset() time.Time {
	return d.sunset
}

func (d Day) Next() calendar.Day {
	next := d.date.Add(1)

	return NewDay(d.calendar, next)
}
