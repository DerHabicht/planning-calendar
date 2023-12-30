package reports

import (
	"fmt"
	"sort"
	"time"

	"github.com/derhabicht/planning-calendar/calendar"
)

type Holiday struct {
	abbreviation string
	fullName     string
	cy1ActDate   time.Time
	cy1ObsDate   time.Time
	cy2ActDate   time.Time
	cy2ObsDate   time.Time
}

func NewHoliday(holiday calendar.Holiday, fy int) Holiday {
	cy1Act, cy1Obs := holiday.Occurs(fy - 1)
	cy2Act, cy2Obs := holiday.Occurs(fy)

	return Holiday{
		abbreviation: holiday.Abbreviation(),
		fullName:     holiday.FullName(),
		cy1ActDate:   cy1Act,
		cy1ObsDate:   cy1Obs,
		cy2ActDate:   cy2Act,
		cy2ObsDate:   cy2Obs,
	}
}

func (h Holiday) LaTeX() string {
	layout := "02 Jan"

	cy1Act := h.cy1ActDate.Format(layout)
	cy2Act := h.cy2ActDate.Format(layout)

	var cy1Obs string
	if h.cy1ObsDate.IsZero() || h.cy1ObsDate.Equal(h.cy1ActDate) {
		cy1Obs = ""
	} else {
		cy1Obs = h.cy1ObsDate.Format(layout)
	}

	var cy2Obs string
	if h.cy2ObsDate.IsZero() || h.cy2ObsDate.Equal(h.cy2ActDate) {
		cy2Obs = ""
	} else {
		cy2Obs = h.cy2ObsDate.Format(layout)
	}

	return fmt.Sprintf("%s & %s & %s & %s & %s & %s \\\\\n", h.abbreviation, h.fullName, cy1Act, cy1Obs, cy2Act, cy2Obs)
}

type HolidayList []Holiday

func NewHolidayList(cal calendar.Calendar) HolidayList {
	holidays := cal.Holidays()
	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Abbreviation() < holidays[j].Abbreviation()
	})

	var hl HolidayList
	for _, holiday := range holidays {
		hl = append(hl, NewHoliday(holiday, cal.FiscalYear()))
	}

	return hl
}

func (hl HolidayList) LaTeX() string {
	latex := ""

	for _, holiday := range hl {
		latex += holiday.LaTeX()
	}

	return latex
}
