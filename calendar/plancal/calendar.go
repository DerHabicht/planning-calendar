package plancal

import (
	"time"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/holidays"
	"github.com/derhabicht/planning-calendar/calendar/solstice_table"
)

const julianPeriodOffest = 4713

type Calendar struct {
	fiscalYear      int
	solsticeTable   calendar.SolsticeTable
	holidayCalendar calendar.HolidayCalendar
}

func NewCalendar(fiscalYear int) *Calendar {
	st := solstice_table.NewSolsticeTable(fiscalYear)
	hc := holidays.NewHolidayCalendar()

	return &Calendar{
		fiscalYear:      fiscalYear,
		solsticeTable:   st,
		holidayCalendar: hc,
	}
}

func (c *Calendar) FiscalYear() int {
	return c.fiscalYear
}

func (c *Calendar) JulianPeriod() int {
	return c.fiscalYear + julianPeriodOffest
}

func (c *Calendar) SolsticeTable() calendar.SolsticeTable {
	return c.solsticeTable
}

func (c *Calendar) HolidayCalendar() calendar.HolidayCalendar {
	return c.holidayCalendar
}

func (c *Calendar) FirstMonth() calendar.Month {
	return NewMonth(c, c.fiscalYear-1, time.October)
}
