package plancal

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/ag7if"
	"github.com/derhabicht/planning-calendar/calendar/doomsday"
	"github.com/derhabicht/planning-calendar/calendar/events"
	"github.com/derhabicht/planning-calendar/internal/config"
)

type Calendar struct {
	fiscalYear          int
	startJulianPeriod   int
	currentCalendarYear int
	currentDoomsday     time.Weekday
	currentMonth        time.Month
	currentWeek         *ag7if.Week
	location            *time.Location
	holidayCalendar     *events.HolidayCalendar
	solsticeTable       events.SolsticeTable
}

func NewCalendar(fiscalYear int) Calendar {
	loc, err := time.LoadLocation(config.GetString("home_location.tz"))
	if err != nil {
		loc = time.UTC
	}

	hc := events.NewHolidayCalendar()

	c := Calendar{
		fiscalYear:          fiscalYear,
		startJulianPeriod:   (fiscalYear - 1) + 4713,
		currentCalendarYear: fiscalYear - 1,
		location:            loc,
		holidayCalendar:     &hc,
		solsticeTable:       events.NewSolsticeTable(fiscalYear),
	}

	c.initMonthAndWeek()

	return c
}

func (c *Calendar) initMonthAndWeek() {
	c.currentCalendarYear = c.fiscalYear - 1
	c.currentDoomsday = doomsday.ComputeDoomsday(c.currentCalendarYear)
	c.currentMonth = time.October
	c.currentWeek = ag7if.NewWeek(date.New(c.currentCalendarYear, time.October, 1), c.location, c.holidayCalendar, c.solsticeTable)
}

func (c *Calendar) CurrentWeek() calendar.Week {
	return c.currentWeek
}

func (c *Calendar) FiscalYear() int {
	return c.fiscalYear
}

func (c *Calendar) SolsticeTable() calendar.SolsticeTable {
	return c.solsticeTable
}

func (c *Calendar) StartingJulianPeriod() int {
	return c.startJulianPeriod
}

func (c *Calendar) NextWeek() *ag7if.Week {
	d := c.currentWeek.Reset()

	c.currentWeek = ag7if.NewWeek(d.Add(7), c.location, c.holidayCalendar, c.solsticeTable)

	return c.currentWeek
}

func (c *Calendar) CurrentMonth() (int, time.Month) {
	return c.currentCalendarYear, c.currentMonth
}

func (c *Calendar) CurrentMonthStr() string {
	return fmt.Sprintf("%s %d", c.currentMonth, c.currentCalendarYear)
}

func (c *Calendar) NextMonth() (string, *ag7if.Week) {
	switch {
	case c.currentMonth == time.December:
		c.currentMonth = time.January
		c.currentCalendarYear += 1
		c.currentDoomsday = doomsday.ComputeDoomsday(c.currentCalendarYear)
	default:
		c.currentMonth += 1
	}

	c.currentWeek = ag7if.NewWeek(date.New(c.currentCalendarYear, c.currentMonth, 1), c.location, c.holidayCalendar, c.solsticeTable)
	return c.CurrentMonthStr(), c.currentWeek
}

func (c *Calendar) Reset() (string, *ag7if.Week) {
	c.initMonthAndWeek()

	return c.CurrentMonthStr(), c.currentWeek
}

func (c *Calendar) Holidays() []calendar.Holiday {
	return c.holidayCalendar.Holidays()
}
