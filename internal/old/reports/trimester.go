package reports

import (
	"fmt"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar/cap"
)

type TrimesterTemplate struct {
	trimester          cap.Trimester
	startDate          date.Date
	miniMonthTemplates []miniMonthTemplate
	template           string
}

func NewTrimesterTemplate(trimester cap.Trimester, template string, miniMonthTemplates []miniMonthTemplate) TrimesterTemplate {
	return TrimesterTemplate{
		trimester:          trimester,
		startDate:          trimester.StartDate(),
		miniMonthTemplates: miniMonthTemplates,
		template:           template,
	}
}

func (t TrimesterTemplate) LaTeX() string {
	template := strings.Replace(t.template, "+T", t.trimester.String(), 1)

	for i, v := range t.miniMonthTemplates {
		key := fmt.Sprintf("+M%dCMD", i+1)
		template = strings.Replace(template, key, v.LaTeXCommand(), 1)
	}

	return template
}
