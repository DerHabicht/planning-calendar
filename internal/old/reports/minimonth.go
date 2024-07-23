package reports

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/doomsday"
)

const miniCalWeekHeaderTemplate = `W & +M & +T & +W & +H & +F & +S & +U \\`

const miniCalMonthTemplate = `\newcommand{+COMMAND}{\fbox{\begin{minipage}{0.24\textwidth}
          \centering
          {\Large\textbf{+MONTH}}\vspace{\baselineskip}
          \begin{tabularx}{\textwidth}{r|rrrrrrr}
              \toprule
 			  +WEEKHEADER
              \midrule
              +W1
              +W2
              +W3
              +W4
              +W5
              +W6
          \end{tabularx}
\end{minipage}}}
`
const miniCalWeekTemplate = `+W & +D1 & +D2 & +D3 & +D4 & +D5 & +D6 & +D7 \\`

func generateTemplates(fy int) []miniMonthTemplate {
	d := date.New(fy-1, time.September, 1)
	endDate := date.New(fy+2, time.June, 1)

	var templates []miniMonthTemplate
	for d.Before(endDate) {
		t := newMiniMonthTemplate(fy, d.Year(), d.Month())

		templates = append(templates, t)

		d = d.AddDate(0, 1, 0)
	}

	return templates
}

type miniMonthTemplate struct {
	month        time.Month
	year         int
	doomsday     time.Weekday
	latexCommand string
}

func newMiniMonthTemplate(fy, year int, month time.Month) miniMonthTemplate {
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
		return miniMonthTemplate{}
	}

	return miniMonthTemplate{
		month:        month,
		year:         year,
		doomsday:     doomsday.ComputeDoomsday(year),
		latexCommand: latexCommand,
	}
}

func (mmt *miniMonthTemplate) generateWeekHeader() string {
	header := miniCalWeekHeaderTemplate
	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		weekdayLetter := calendar.WeekdayLetters[weekday]

		repl := weekdayLetter
		if weekday == mmt.doomsday {
			repl = fmt.Sprintf("\\underline{%s}", weekdayLetter)
		}

		header = strings.Replace(header, fmt.Sprintf("+%s", weekdayLetter), repl, 1)
	}

	return header
}

func (mmt *miniMonthTemplate) fillWeeks(template string, start date.Date) string {
	d := start
	startMonth := d.Month()
	for week := 1; week <= 6; week++ {
		wkTemplate := miniCalWeekTemplate
		if d.Month() == startMonth {
			_, isoWeek := d.ISOWeek()
			wkTemplate = strings.Replace(wkTemplate, "+W", strconv.Itoa(isoWeek), 1)
		} else {
			wkTemplate = strings.Replace(wkTemplate, "+W", "", 1)
		}
		for weekDay := 1; weekDay <= 7; weekDay++ {
			if (int(d.Weekday()) != weekDay%7) || d.Month() != startMonth {
				wkTemplate = strings.Replace(wkTemplate, fmt.Sprintf("+D%d", weekDay), "", 1)
				continue
			}
			wkTemplate = strings.Replace(wkTemplate, fmt.Sprintf("+D%d", weekDay), strconv.Itoa(d.Day()), 1)
			d = d.Add(1)
		}

		template = strings.Replace(template, fmt.Sprintf("+W%d", week), wkTemplate, 1)
	}

	return template
}

func (mmt *miniMonthTemplate) MonthStr() string {
	return fmt.Sprintf("%s %d", mmt.month, mmt.year)
}

func (mmt *miniMonthTemplate) LaTeXCommand() string {
	return mmt.latexCommand
}

func (mmt *miniMonthTemplate) LaTeX() string {
	latex := miniCalMonthTemplate
	latex = strings.Replace(latex, "+COMMAND", mmt.latexCommand, 1)
	latex = strings.Replace(latex, "+MONTH", fmt.Sprintf("%s %d", mmt.month, mmt.year), 1)
	latex = strings.Replace(latex, "+WEEKHEADER", mmt.generateWeekHeader(), 1)
	latex = mmt.fillWeeks(latex, date.New(mmt.year, mmt.month, 1))

	return latex
}
