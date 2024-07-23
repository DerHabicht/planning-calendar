package events

import (
	"time"

	"github.com/rickar/cal/v2"

	"github.com/derhabicht/planning-calendar/calendar"
)

type Holiday struct {
	abbreviation string
	holiday      *cal.Holiday
}

func (h Holiday) Abbreviation() string {
	return h.abbreviation
}

func (h Holiday) FullName() string {
	return h.holiday.Name
}

func (h Holiday) Occurs(year int) (time.Time, time.Time) {
	return h.holiday.Calc(year)
}

func (h Holiday) CalHoliday() *cal.Holiday {
	return h.holiday
}

type HolidayCalendar struct {
	holidays map[string]calendar.Holiday
	calendar *cal.Calendar
}

func NewHolidayCalendar() HolidayCalendar {
	c := HolidayCalendar{
		holidays: make(map[string]calendar.Holiday),
		calendar: &cal.Calendar{},
	}

	for abbv, h := range Ag7ifHolidays {
		c.AddHoliday(Holiday{abbv, h})
	}

	return c
}

func (h *HolidayCalendar) AddHoliday(holiday calendar.Holiday) {
	h.holidays[holiday.FullName()] = holiday
	h.calendar.AddHoliday(holiday.CalHoliday())
}

func (h *HolidayCalendar) IsHoliday(date time.Time) (bool, bool, calendar.Holiday) {
	act, obs, calHoliday := h.calendar.IsHoliday(date)
	if calHoliday != nil {
		holiday := h.holidays[calHoliday.Name]
		return act, obs, holiday
	}

	return false, false, nil
}

func (h *HolidayCalendar) Holidays() []calendar.Holiday {
	holidays := make([]calendar.Holiday, 0)

	for _, holiday := range h.holidays {
		holidays = append(holidays, holiday)
	}

	return holidays
}
