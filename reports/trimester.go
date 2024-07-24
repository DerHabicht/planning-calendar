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
	minimonths []Minimonth
}

func NewTrimester(trimester calendar.Trimester, minimonths map[string]Minimonth) Trimester {
	var mm []Minimonth

	d := trimester.StartDate()
	for i := 0; i < trimesterMonthCount; i++ {
		mm = append(mm, minimonths[d.Format(minimonthKeyLayout)])
		d.AddDate(0, 1, 0)
	}

	return Trimester{
		trimester:  trimester,
		minimonths: mm,
	}
}

func (t *Trimester) LaTeX() string {
	latex := templates.TrimesterTemplate

	latex = strings.Replace(latex, templates.FullTrimester, t.trimester.Full(), 1)

	d := date.New(t.trimester.StartDate().Year(), t.trimester.StartDate().Month(), 1)
	for i, mm := range t.minimonths {
		latex = strings.Replace(latex, fmt.Sprintf(templates.MinimonthMacro, i+1), mm.LatexCommand(), 1)
		d = d.AddDate(0, 1, 0)
	}

	return latex
}
