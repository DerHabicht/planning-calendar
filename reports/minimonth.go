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

const minimonthWeekCount = 6
const minimonthKeyLayout = `2006-01`

type Minimonth struct {
	month        time.Month
	year         int
	doomsday     time.Weekday
	latexCommand string
}

func NewMinimonthList(calendar calendar.Calendar) map[string]Minimonth {
	minimonths := make(map[string]Minimonth)

	endDate := date.New(calendar.FiscalYear()+2, time.June, 1)
	d := date.New(calendar.FiscalYear()-1, time.September, 1)
	for d.Before(endDate) {
		m := NewMinimonth(calendar.FiscalYear(), d.Year(), d.Month())

		minimonths[d.Format(minimonthKeyLayout)] = m

		d = d.AddDate(0, 1, 0)
	}

	return minimonths
}

func NewMinimonth(fy, year int, month time.Month) Minimonth {
	latexCommand := `\mini`

	// Work out the LaTeX command and template keys for the current year/month combo.
	if (year == fy-1) || (year == fy && month == time.January) {
		// Case 1: Oct-Dec of FY-1 and Jan of FY
		latexCommand = fmt.Sprintf("%sfirst%s", latexCommand, strings.ToLower(month.String()))
	} else if year == fy && (month >= time.February && month <= time.August) {
		// Case 2: Feb-Aug of FY
		latexCommand = fmt.Sprintf("%s%s", latexCommand, strings.ToLower(month.String()))
	} else if (year == fy && month >= time.September) || (year == fy+1 && month == time.January) {
		// Case 3: Sep-Dec of FY, and Jan of FY+1
		latexCommand = fmt.Sprintf("%ssecond%s", latexCommand, strings.ToLower(month.String()))
	} else if year == fy+1 && (month >= time.February && month <= time.May) {
		// Case 4: Feb-May of FY+1
		latexCommand = fmt.Sprintf("%sfirstext%s", latexCommand, strings.ToLower(month.String()))
	} else if (year == fy+1 && month >= time.June) || (year == fy+2 && month == time.January) {
		// Case 5: Jun-Dec of FY+1, and Jan of FY+2
		latexCommand = fmt.Sprintf("%sext%s", latexCommand, strings.ToLower(month.String()))
	} else if year == fy+2 && (month >= time.February && month <= time.May) {
		// Case 6: Feb-May of FY+2
		latexCommand = fmt.Sprintf("%ssecondext%s", latexCommand, strings.ToLower(month.String()))
	} else {
		return Minimonth{}
	}

	return Minimonth{
		month:        month,
		year:         year,
		doomsday:     doomsday.ComputeDoomsday(year),
		latexCommand: latexCommand,
	}
}

func (mm *Minimonth) generateWeekHeader(latex string) string {
	templ := templates.MinimonthWeekHeaderTemplate

	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		letter := calendar.WeekdayLetter(weekday)

		repl := letter
		if weekday == mm.doomsday {
			repl = fmt.Sprintf(`\underline{%s}`, repl)
		}

		templ = strings.Replace(templ, fmt.Sprintf("+%s", letter), repl, 1)
	}

	latex = strings.Replace(latex, templates.WeekHeader, templ, 1)

	return latex
}

func (mm *Minimonth) generateWeeks(latex string) string {
	d := date.New(mm.year, mm.month, 1)
	for week := 1; week <= minimonthWeekCount; week++ {
		templ := templates.MinimonthWeekTemplate

		_, isoWeek := d.ISOWeek()
		templ = strings.Replace(templ, templates.MinimonthWeekNumber, strconv.Itoa(isoWeek), 1)

		for weekDay := 1; weekDay <= 7; weekDay++ {
			if (int(d.Weekday()) != weekDay%7) || d.Month() != mm.month {
				templ = strings.Replace(templ, fmt.Sprintf(templates.MinimonthDay, weekDay), "", 1)
				continue
			}
			templ = strings.Replace(templ, fmt.Sprintf(templates.MinimonthDay, weekDay), strconv.Itoa(d.Day()), 1)
			d = d.Add(1)
		}

		latex = strings.Replace(latex, fmt.Sprintf(templates.MinimonthWeek, week), templ, 1)
	}

	return latex
}

func (mm *Minimonth) LatexCommand() string {
	return mm.latexCommand
}

func (mm *Minimonth) LaTeX() string {
	latex := templates.MinimonthTemplate

	latex = strings.Replace(latex, templates.MinimonthCommand, mm.latexCommand, 1)
	latex = strings.Replace(latex, templates.MonthName, fmt.Sprintf("%s %d", mm.month, mm.year), 1)
	latex = mm.generateWeekHeader(latex)
	latex = mm.generateWeeks(latex)

	return latex
}
