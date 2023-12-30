package calendar

import (
	"time"

	"github.com/rickar/cal/v2"
)

type Holiday struct {
	abbreviation string
	showInList   bool
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

// CalHoliday returns a pointer to the underlying cal.Holiday struct.
// WARNING: This method is intended to facilitate the HolidayCalendar struct and should *never* be called outside
// of this calendar package.
func (h Holiday) CalHoliday() *cal.Holiday {
	return h.holiday
}

type HolidayCalendar struct {
	holidays map[string]Holiday
	calendar *cal.Calendar
}

func NewHolidayCalendar() HolidayCalendar {
	c := HolidayCalendar{
		holidays: make(map[string]Holiday),
		calendar: &cal.Calendar{},
	}

	for abbv, h := range Ag7ifHolidays {
		c.AddHoliday(Holiday{abbv, true, h})
	}

	return c
}

func (h *HolidayCalendar) AddHoliday(holiday Holiday) {
	h.holidays[holiday.FullName()] = holiday
	h.calendar.AddHoliday(holiday.CalHoliday())
}

func (h *HolidayCalendar) IsHoliday(date time.Time) (bool, bool, *Holiday) {
	act, obs, calHoliday := h.calendar.IsHoliday(date)
	if calHoliday != nil {
		holiday := h.holidays[calHoliday.Name]
		return act, obs, &holiday
	}

	return false, false, nil
}
