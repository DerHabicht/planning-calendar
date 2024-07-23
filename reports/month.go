package reports

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/doomsday"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

type Month struct {
	calendar calendar.Calendar
	month    calendar.Month
}

func NewMonth(calendar calendar.Calendar, month calendar.Month) Month {

	return Month{
		calendar: calendar,
		month:    month,
	}
}

func (m *Month) generateWeekdayHeader(latex string) string {
	header := templates.MonthWeekdayHeaderTemplate

	year := m.month.Year()
	dd := doomsday.ComputeDoomsday(year)

	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		abbv := strings.ToUpper(weekday.String())[:3]

		repl := abbv
		if weekday == dd {
			repl = fmt.Sprintf("<%s>", abbv)
		}

		header = strings.Replace(header, fmt.Sprintf("+%s", abbv), repl, 1)
	}

	latex = strings.Replace(latex, "+WEEKDAYS", header, 1)

	return latex
}

func (m *Month) generateWeekData(latex string) string {
	return latex
}

func (m *Month) generateDayData(latex string) string {
	const timeFormat = `1504`

	d := m.month.FirstWeek().Monday()
	for i := 1; i <= 42; i++ {
		day := templates.MonthDayTemplate

		dayStr := strconv.Itoa(d.Date().Day())
		if i == 1 || d.Date().Day() == 1 {
			dayStr = strings.ToUpper(d.Date().Month().String())[:3] + " " + dayStr
		}

		solstice := d.IsSolstice()
		if solstice != calendar.NoSolstice {
			dayStr = fmt.Sprintf(`%s\hfill{}%s`, solstice.LaTeX(), dayStr)
		}

		day = strings.Replace(day, "+DY", dayStr, 1)

		act, obs, holiday := d.IsHoliday()
		if act {
			day = strings.Replace(day, "+HD", holiday.String(), 1)
		} else if obs {
			day = strings.Replace(day, "+HD", fmt.Sprintf("%s*", holiday), 1)
		} else {
			day = strings.Replace(day, "+HD", "", 1)
		}

		day = strings.Replace(day, "+FD", d.ISODate(), 1)
		day = strings.Replace(day, "+YD", strconv.Itoa(d.OrdinalDay()), 1)
		day = strings.Replace(day, "+SR", d.Sunrise().Format(timeFormat), 1)
		day = strings.Replace(day, "+MJD", strconv.Itoa(d.MJD()), 1)
		day = strings.Replace(day, "+SS", d.Sunset().Format(timeFormat), 1)

		latex = strings.Replace(latex, fmt.Sprintf("+D%0d", i), day, 1)
	}

	return latex
}

func (m *Month) LaTeX() string {
	latex := templates.MonthTemplate

	latex = m.generateWeekdayHeader(latex)
	latex = m.generateWeekData(latex)
	latex = m.generateDayData(latex)

	return latex
}
