package reports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

const quarterWeekCount = 13
const quarterMonthCount = 3

type Quarter struct {
	quarter    calendar.Quarter
	minimonths []Minimonth
}

func NewQuarter(quarter calendar.Quarter, minimonths map[string]Minimonth) Quarter {
	var mm []Minimonth

	d := quarter.StartDate()
	for i := 0; i < quarterMonthCount; i++ {
		mm = append(mm, minimonths[d.Format(minimonthKeyLayout)])
		d.AddDate(0, 1, 0)
	}

	return Quarter{
		quarter:    quarter,
		minimonths: mm,
	}
}

func (q *Quarter) generateOKRHeader(latex string) string {
	hdr := templates.OKRHeaderTemplate

	w := q.quarter.FirstWeek()
	for i := 1; i <= quarterWeekCount; i++ {
		_, wk, _ := w.ISOWeek()
		hdr = strings.Replace(hdr, fmt.Sprintf(templates.OKRHeaderWeekNumber, i), strconv.Itoa(wk), 1)
	}

	latex = strings.Replace(latex, templates.OKRHeader, hdr, -1)

	return latex
}

func (q *Quarter) LaTeX() string {
	latex := templates.QuarterTemplate

	latex = strings.Replace(latex, templates.FullFiscalQuarter, q.quarter.FirstWeek().FiscalQuarter().Full(), 1)
	latex = strings.Replace(latex, templates.FullCalendarQuarter, q.quarter.FirstWeek().Quarter().Full(), 1)

	d := date.New(q.quarter.StartDate().Year(), q.quarter.StartDate().Month(), 1)
	for i, mm := range q.minimonths {
		latex = strings.Replace(latex, fmt.Sprintf(templates.MinimonthMacro, i+1), mm.LatexCommand(), 1)
		d = d.AddDate(0, 1, 0)
	}

	latex = q.generateOKRHeader(latex)

	return latex
}
