package reports

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/internal/config"
)

type Calendar struct {
	calendar           calendar.Calendar
	solsticeTable      SolsticeTable
	miniMonthTemplates []miniMonthTemplate
	calendarTemplate   string
	trimesterTemplate  string
	quarterTemplate    string
	monthTemplate      string
}

func NewCalendar(cal calendar.Calendar) (Calendar, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}

	raw, err := os.ReadFile(filepath.Join(cfgDir, "assets", "calendar.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	calendarTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "trimester.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	trimesterTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "quarter.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	quarterTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "month.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	monthTemplate := string(raw)

	miniMonthTemplates := generateTemplates(cal.FiscalYear())

	return Calendar{
		calendar:           cal,
		solsticeTable:      NewSolsticeTable(cal.SolsticeTable()),
		miniMonthTemplates: miniMonthTemplates,
		calendarTemplate:   calendarTemplate,
		trimesterTemplate:  trimesterTemplate,
		quarterTemplate:    quarterTemplate,
		monthTemplate:      monthTemplate,
	}, nil
}

func (ct Calendar) generateDoomsdayContextTable() string {
	latex := `\begin{tabular}{rc}
\toprule
\textbf{Year} & \textbf{Doomsday} \\
\midrule
`

	for y := ct.calendar.FiscalYear() - 2; y < ct.calendar.FiscalYear()+3; y++ {
		doomsday := calendar.WeekdayLetters[calendar.ComputeDoomsday(y)]

		if y == ct.calendar.FiscalYear() {
			latex += fmt.Sprintf("\\textbf{%d} & \\textbf{%s} \\\\\n", y, doomsday)
		} else {
			latex += fmt.Sprintf("%d & %s \\\\\n", y, doomsday)
		}
	}

	latex += `\bottomrule\end{tabular}`

	return latex
}

func (ct Calendar) LaTeX() string {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		panic(err)
	}
	coverLogoPath := filepath.Join(cfgDir, "assets", config.GetString("cover_logo"))

	holidayList := NewHolidayList(ct.calendar)

	latex := ct.calendarTemplate

	latex = strings.Replace(latex, "+PIC", coverLogoPath, 1)
	latex = strings.Replace(latex, "+CAL_START", fmt.Sprintf("October %d", ct.calendar.FiscalYear()-1), 1)
	latex = strings.Replace(latex, "+CAL_END", fmt.Sprintf("December %d", ct.calendar.FiscalYear()), 1)
	latex = strings.Replace(latex, "+JP_START", fmt.Sprintf("%d", ct.calendar.StartingJulianPeriod()), 1)
	latex = strings.Replace(latex, "+JP_END", fmt.Sprintf("%d", ct.calendar.StartingJulianPeriod()+1), 1)
	latex = strings.Replace(latex, "+LCD", config.GetString("lunar_calibration_date"), 1)
	latex = strings.Replace(latex, "+CY1", fmt.Sprintf("%d", ct.calendar.FiscalYear()-1), 2)
	latex = strings.Replace(latex, "+CY2", fmt.Sprintf("%d", ct.calendar.FiscalYear()), 2)
	latex = strings.Replace(latex, "+ABBVS", holidayList.LaTeX(), 1)
	latex = strings.Replace(latex, "+DOOMSDAYS", ct.generateDoomsdayContextTable(), 1)
	latex = strings.Replace(latex, "+SOLSTICES", ct.solsticeTable.LaTeX(), 1)

	for _, v := range ct.miniMonthTemplates {
		latex = strings.Replace(latex, v.TemplateKey(), v.LaTeX(), 1)
	}

	tri := calendar.FyT1
	for trimester := 0; trimester < 4; trimester++ {
		startMonthIdx := (trimester * 4) + 1
		endMonthIdx := (trimester * 4) + 5
		tt := NewTrimesterTemplate(tri, ct.calendar.FiscalYear(), ct.trimesterTemplate, ct.miniMonthTemplates[startMonthIdx:endMonthIdx])

		latex = strings.Replace(latex, fmt.Sprintf("+T%d", trimester+1), tt.LaTeX(), 1)

		_, tri = tri.NextTrimester(ct.calendar.FiscalYear())
	}

	fyQtr := calendar.FyQ1
	for quarter := 0; quarter < 5; quarter++ {
		startMonthIdx := (quarter * 3) + 1
		endMonthIdx := (quarter * 3) + 4
		qt := NewQuarterTemplate(fyQtr, ct.calendar.FiscalYear(), ct.quarterTemplate, ct.miniMonthTemplates[startMonthIdx:endMonthIdx])

		latex = strings.Replace(latex, fmt.Sprintf("+Q%d", quarter+1), qt.LaTeX(), 1)

		_, fyQtr = fyQtr.NextQuarter(ct.calendar.FiscalYear())
	}

	for month := 1; month <= 15; month++ {
		mt := NewMonthTemplate(ct.calendar, ct.monthTemplate)

		latex = strings.Replace(latex, fmt.Sprintf("+M%02d", month), mt.LaTeX(), 1)
		latex = strings.Replace(latex, "+PREVCMD", ct.miniMonthTemplates[month-1].LaTeXCommand(), 1)
		latex = strings.Replace(latex, "+NEXTCMD", ct.miniMonthTemplates[month+1].LaTeXCommand(), 1)

		ct.calendar.NextMonth()
	}

	return latex
}
