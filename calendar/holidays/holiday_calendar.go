package holidays

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/rickar/cal/v2"

	"github.com/derhabicht/planning-calendar/calendar"
)

type Holiday struct {
	abbreviation string
	holiday      *cal.Holiday
}

func NewHoliday(abbv string, holiday *cal.Holiday) Holiday {
	return Holiday{
		abbreviation: abbv,
		holiday:      holiday,
	}
}

func (h Holiday) Occurs(year int) (date.Date, date.Date) {
	act, obs := h.holiday.Calc(year)

	return date.New(act.Year(), act.Month(), act.Day()), date.New(obs.Year(), obs.Month(), obs.Day())
}

func (h Holiday) String() string {
	return h.abbreviation
}

func (h Holiday) FullName() string {
	return h.holiday.Name
}

type HolidayCalendar struct {
	calendar *cal.Calendar
	holidays map[string]Holiday
	loc      *time.Location
}

func NewHolidayCalendar() *HolidayCalendar {
	holidayCalendar := new(cal.Calendar)
	holidays := make(map[string]Holiday)
	loc := calendar.GetLocation()

	for k, v := range Ag7ifHolidays {
		holidayCalendar.AddHoliday(v)
		h := NewHoliday(k, v)
		holidays[h.FullName()] = h
	}

	return &HolidayCalendar{
		calendar: holidayCalendar,
		holidays: holidays,
		loc:      loc,
	}
}

func (h *HolidayCalendar) IsHoliday(date date.Date) (bool, bool, calendar.Holiday) {
	act, obs, calHoliday := h.calendar.IsHoliday(date.In(h.loc))

	if calHoliday != nil {
		return act, obs, h.holidays[calHoliday.Name]
	}

	return false, false, nil
}

func (h *HolidayCalendar) Holidays() []calendar.Holiday {
	var holidays []calendar.Holiday

	for _, v := range h.holidays {
		holidays = append(holidays, v)
	}

	return holidays
}
