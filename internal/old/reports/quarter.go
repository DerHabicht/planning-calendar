package reports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar/ag7if"
	"github.com/derhabicht/planning-calendar/calendar/cap"
)

const okrHeaderTemplate = `Objectives & Key Results & +W01 & +W02 & +W03 & +W04 & +W05 & +W06 & +W07 & +W08 & +W09 & +W10 & +W11 & +W12 & +W13`

type QuarterTemplate struct {
	fiscalQuarter      cap.Quarter
	calQuarter         ag7if.Quarter
	startDate          date.Date
	miniMonthTemplates []miniMonthTemplate
	template           string
}

func NewQuarterTemplate(quarter cap.Quarter, template string, miniMonthTemplates []miniMonthTemplate) QuarterTemplate {
	calQuarter := quarter.CyQuarter()

	return QuarterTemplate{
		fiscalQuarter:      quarter,
		calQuarter:         calQuarter,
		startDate:          quarter.StartDate(),
		miniMonthTemplates: miniMonthTemplates,
		template:           template,
	}
}

func (q QuarterTemplate) setOKRHeader(template string) string {
	header := okrHeaderTemplate
	week := q.calQuarter.Quarter.StartWeek()
	for w := 1; w <= 13; w++ {
		header = strings.Replace(header, fmt.Sprintf("+W%02d", w), strconv.Itoa(week), 1)
		week++
	}

	return strings.Replace(template, "+OKRHDR", header, -1)
}

func (q QuarterTemplate) LaTeX() string {
	template := strings.Replace(q.template, "+CYQ", q.calQuarter.String(), 1)
	template = strings.Replace(template, "+FYQ", q.fiscalQuarter.String(), 1)
	template = q.setOKRHeader(template)

	for i, v := range q.miniMonthTemplates {
		key := fmt.Sprintf("+M%dCMD", i+1)
		template = strings.Replace(template, key, v.LaTeXCommand(), 1)
	}

	return template
}
