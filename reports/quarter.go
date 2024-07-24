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
	minimonths map[date.Date]Minimonth
}

func NewQuarter(quarter calendar.Quarter, minimonths map[date.Date]Minimonth) Quarter {
	return Quarter{
		quarter:    quarter,
		minimonths: minimonths,
	}
}

func (q *Quarter) generateOKRHeader(latex string) string {
	hdr := templates.OKRHeaderTemplate

	w := q.quarter.FirstWeek()
	for i := 1; i <= quarterWeekCount; i++ {
		_, wk, _ := w.ISOWeek()
		hdr = strings.Replace(hdr, fmt.Sprintf("+W%0d", i), strconv.Itoa(wk), 1)
	}

	latex = strings.Replace(latex, "+OKR_HDR", hdr, 1)

	return latex
}

func (q *Quarter) LaTeX() string {
	latex := templates.QuarterTemplate

	latex = strings.Replace(latex, "+FYQ", q.quarter.FirstWeek().FiscalQuarter().Full(), 1)
	latex = strings.Replace(latex, "+CYQ", q.quarter.FirstWeek().Quarter().Full(), 1)

	d := date.New(q.quarter.StartDate().Year(), q.quarter.StartDate().Month(), 1)
	for i := 1; i <= quarterMonthCount; i++ {
		mm := q.minimonths[d]
		latex = strings.Replace(latex, fmt.Sprintf("+M%dCMD", i), mm.LatexCommand(), 1)
		d = d.AddDate(0, 1, 0)
	}

	latex = q.generateOKRHeader(latex)

	return latex
}
