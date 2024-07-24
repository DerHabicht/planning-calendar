package reports

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/doomsday"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

const monthWeekCount = 6
const dayCount = monthWeekCount * 7

type Month struct {
	calendar  calendar.Calendar
	month     calendar.Month
	prevMonth Minimonth
	nextMonth Minimonth
}

func NewMonth(calendar calendar.Calendar, month calendar.Month, minimonths map[string]Minimonth) Month {
	m := date.New(month.Year(), month.Month(), 1)
	p := m.AddDate(0, -1, 0)
	n := m.AddDate(0, 1, 0)

	return Month{
		calendar:  calendar,
		month:     month,
		prevMonth: minimonths[p.Format(minimonthKeyLayout)],
		nextMonth: minimonths[n.Format(minimonthKeyLayout)],
	}
}

func (m *Month) generateMinimonths(latex string) string {
	latex = strings.Replace(latex, templates.MinimonthPrevious, m.prevMonth.LatexCommand(), 1)
	latex = strings.Replace(latex, templates.MinimonthNext, m.nextMonth.LatexCommand(), 1)

	return latex
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

	latex = strings.Replace(latex, templates.WeekdayHeader, header, 1)

	return latex
}

func (m *Month) generateWeekData(latex string) string {
	week := m.month.FirstWeek()
	for i := 1; i <= monthWeekCount; i++ {
		_, fyWeek := week.FyWeek()
		_, cyWeek, card := week.ISOWeek()

		latex = strings.Replace(latex, fmt.Sprintf(templates.FiscalTrimester, i), week.Trimester().Short(), 1)
		latex = strings.Replace(latex, fmt.Sprintf(templates.FiscalQuarter, i), week.FiscalQuarter().Short(), 1)
		latex = strings.Replace(latex, fmt.Sprintf(templates.FiscalWeek, i), strconv.Itoa(fyWeek), 1)
		latex = strings.Replace(latex, fmt.Sprintf(templates.CalendarQuarter, i), week.Quarter().Short(), 1)
		latex = strings.Replace(latex, fmt.Sprintf(templates.Sprint, i), week.Sprint().Short(), 1)
		latex = strings.Replace(latex, fmt.Sprintf(templates.ISOWeek, i), fmt.Sprintf("%sW%0d", card.LaTeX(), cyWeek), 1)

		week = week.Next()
	}

	return latex
}

func (m *Month) generateDayData(latex string) string {
	const timeFormat = `1504`

	d := m.month.FirstWeek().Monday()
	for i := 1; i <= dayCount; i++ {
		day := templates.MonthDayTemplate

		dayStr := strconv.Itoa(d.Date().Day())
		if i == 1 || d.Date().Day() == 1 {
			dayStr = strings.ToUpper(d.Date().Month().String())[:3] + " " + dayStr
		}

		solstice := d.IsSolstice()
		if solstice != calendar.NoSolstice {
			dayStr = fmt.Sprintf(`%s\hfill{}%s`, solstice.LaTeX(), dayStr)
		}

		day = strings.Replace(day, templates.MonthDay, dayStr, 1)

		act, obs, holiday := d.IsHoliday()
		if act {
			day = strings.Replace(day, templates.Holiday, holiday.String(), 1)
		} else if obs {
			day = strings.Replace(day, templates.Holiday, fmt.Sprintf("%s*", holiday), 1)
		} else {
			day = strings.Replace(day, templates.Holiday, "", 1)
		}

		day = strings.Replace(day, templates.MoonPhase, d.MoonPhase().LaTeX(), 1)
		day = strings.Replace(day, templates.OrdinalDay, strconv.Itoa(d.OrdinalDay()), 1)
		day = strings.Replace(day, templates.SunriseTime, d.Sunrise().Format(timeFormat), 1)
		day = strings.Replace(day, templates.MJD, strconv.Itoa(d.MJD()), 1)
		day = strings.Replace(day, templates.SunsetTime, d.Sunset().Format(timeFormat), 1)

		latex = strings.Replace(latex, fmt.Sprintf(templates.MonthDayData, i), day, 1)

		d = d.Next()
	}

	return latex
}

func (m *Month) LaTeX() string {
	latex := templates.MonthTemplate

	latex = strings.Replace(latex, templates.MonthNameFull, m.month.String(), 1)
	latex = m.generateMinimonths(latex)
	latex = m.generateWeekdayHeader(latex)
	latex = m.generateWeekData(latex)
	latex = m.generateDayData(latex)

	return latex
}
