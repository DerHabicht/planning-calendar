package reports

import (
	"slices"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

type HolidayData struct {
	abbreviation string
	fullName     string
	cy1act       date.Date
	cy1obs       date.Date
	cy2act       date.Date
	cy2obs       date.Date
}

func NewHolidayData(fiscalYear int, holiday calendar.Holiday) HolidayData {
	cy1act, cy1obs := holiday.Occurs(fiscalYear - 1)
	cy2act, cy2obs := holiday.Occurs(fiscalYear)

	return HolidayData{
		abbreviation: holiday.String(),
		fullName:     holiday.FullName(),
		cy1act:       cy1act,
		cy1obs:       cy1obs,
		cy2act:       cy2act,
		cy2obs:       cy2obs,
	}
}

func (hd HolidayData) Occurrence() date.Date {
	return hd.cy1act
}

func (hd HolidayData) String() string {
	return hd.abbreviation
}

func (hd HolidayData) LaTeX() string {
	const layout = `02 Jan`
	latex := templates.HolidayAbbvRowTemplate

	latex = strings.Replace(latex, templates.HolidayAbbreviation, hd.abbreviation, 1)
	latex = strings.Replace(latex, templates.HolidayName, hd.fullName, 1)
	latex = strings.Replace(latex, templates.HolidayActual1, hd.cy1act.Format(layout), 1)
	latex = strings.Replace(latex, templates.HolidayActual2, hd.cy2act.Format(layout), 1)

	if hd.cy1act.Equal(hd.cy1obs) {
		latex = strings.Replace(latex, templates.HolidayObserved1, "", 1)
	} else {
		latex = strings.Replace(latex, templates.HolidayObserved1, hd.cy1act.Format(layout), 1)
	}

	if hd.cy2act.Equal(hd.cy2obs) {
		latex = strings.Replace(latex, templates.HolidayObserved2, "", 2)
	} else {
		latex = strings.Replace(latex, templates.HolidayObserved2, hd.cy2act.Format(layout), 2)
	}

	return latex
}

type HolidayTables struct {
	holidays []HolidayData
}

func NewHolidayTables(cal calendar.HolidayCalendar, fy int) HolidayTables {
	holidays := cal.Holidays()

	var data []HolidayData
	for _, holiday := range holidays {
		d := NewHolidayData(fy, holiday)
		data = append(data, d)
	}

	return HolidayTables{
		holidays: data,
	}
}

func (ht HolidayTables) TableByOccurrence(latex string) string {
	slices.SortFunc(ht.holidays, func(a, b HolidayData) int {
		if a.Occurrence().Before(b.Occurrence()) {
			return -1
		} else if b.Occurrence().Before(a.Occurrence()) {
			return 1
		}

		return strings.Compare(a.String(), b.String())
	})

	table := ""
	for _, holiday := range ht.holidays {
		table += holiday.LaTeX()
	}

	latex = strings.Replace(latex, templates.HolidayTableByOccurrence, table, 1)

	return latex
}

func (ht HolidayTables) TableByAbbreviation(latex string) string {
	slices.SortFunc(ht.holidays, func(a, b HolidayData) int {
		return strings.Compare(a.String(), b.String())
	})

	table := ""
	for _, holiday := range ht.holidays {
		table += holiday.LaTeX()
	}

	latex = strings.Replace(latex, templates.HolidayTableByAbbreviation, table, 1)

	return latex
}
