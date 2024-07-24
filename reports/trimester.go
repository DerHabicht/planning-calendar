package reports

import (
	"fmt"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

const trimesterMonthCount = 4

type Trimester struct {
	trimester  calendar.Trimester
	minimonths map[date.Date]Minimonth
}

func NewTrimester(trimester calendar.Trimester, minimonths map[date.Date]Minimonth) Trimester {
	return Trimester{
		trimester:  trimester,
		minimonths: minimonths,
	}
}

func (t *Trimester) LaTeX() string {
	latex := templates.TrimesterTemplate

	latex = strings.Replace(latex, "+T", t.trimester.Full(), 1)

	d := date.New(t.trimester.StartDate().Year(), t.trimester.StartDate().Month(), 1)
	for i := 1; i <= trimesterMonthCount; i++ {
		mm := t.minimonths[d]
		latex = strings.Replace(latex, fmt.Sprintf("+M%dCMD", i), mm.LatexCommand(), 1)
		d = d.AddDate(0, 1, 0)
	}

	return latex
}
