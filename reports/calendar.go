package reports

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/calendar/doomsday"
	"github.com/derhabicht/planning-calendar/internal/config"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

const calendarMonthCount = 15
const trimesterCount = 4 // 15/4 = 4.75, so we round up
const quarterCount = 5   // 15/3 = 5

type Calendar struct {
	cfgDir     string
	calendar   calendar.Calendar
	minimonths map[date.Date]Minimonth
}

func NewCalendar(cal calendar.Calendar) *Calendar {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		panic(errors.WithStack(err))
	}

	minimonths := NewMinimonthList(cal)

	return &Calendar{
		cfgDir:     cfgDir,
		calendar:   cal,
		minimonths: minimonths,
	}
}

func (c *Calendar) fillCalParams(latex string) string {
	//	Set the Lunar Calibration Date that the Tikz uses to calculate and then draw the phase of the moon
	latex = strings.Replace(latex, "+LCD", config.GetString("lunar_calibration_date"), 1)
	//	Set the full name and year of the first full month page in this calendar
	latex = strings.Replace(latex, "+CAL_START", fmt.Sprintf("October %d", c.calendar.FiscalYear()-1), 1)
	//	Set the full name and year of the last full month page in this calendar
	latex = strings.Replace(latex, "+CAL_END", fmt.Sprintf("December %d", c.calendar.FiscalYear()), 1)
	//	Set the starting year of this calendar, expressed as the year of Julian Period A
	latex = strings.Replace(latex, "+JP_START", strconv.Itoa(c.calendar.JulianPeriod()-1), 1)
	//	Set the ending year of this calendar, expressed as the year of Julian Period A
	latex = strings.Replace(latex, "+JP_END", strconv.Itoa(c.calendar.JulianPeriod()), 1)
	//	Set the picture to typeset on the title page of the calendar
	latex = strings.Replace(latex, "+PIC", filepath.Join(c.cfgDir, "assets", config.GetString("cover_logo")), 1)
	//	Set the first calendar year covered in this calendar (i.e. FY-1)
	latex = strings.Replace(latex, "+CY1", strconv.Itoa(c.calendar.FiscalYear()-1), 1)
	//	Set the second calendar year covered in this calendar (i.e. FY)
	latex = strings.Replace(latex, "+CY2", strconv.Itoa(c.calendar.FiscalYear()), 1)

	//	Set the color to use for the title box outline on the calendar's title page
	if c.calendar.FiscalYear()%2 == 0 {
		latex = strings.Replace(latex, "+TITLE_COLOR", "blue", 1)
	} else {
		latex = strings.Replace(latex, "+TITLE_COLOR", "red", 1)
	}

	return latex
}

func (c *Calendar) generateDoomsdayTable(latex string) string {
	table := templates.DoomsdayTableTemplate

	var rows string
	for year := c.calendar.FiscalYear() - 2; year <= c.calendar.FiscalYear()+2; year++ {
		dd := doomsday.ComputeDoomsday(year)
		row := templates.DoomsdayTableRowTemplate
		if year == c.calendar.FiscalYear() {
			row = strings.Replace(row, "+Y", fmt.Sprintf(`\textbf{%d}`, year), 1)
			row = strings.Replace(row, "+DD", fmt.Sprintf(`\textbf{%d}`, dd), 1)
		} else {
			row = strings.Replace(row, "+Y", strconv.Itoa(year), 1)
			row = strings.Replace(row, "+DD", calendar.WeekdayLetter(dd), 1)
		}

		rows += row
	}

	table = strings.Replace(table, "+DD_TABLE_ROWS", rows, 1)
	latex = strings.Replace(latex, "+DOOMSDAYS", table, 1)

	return latex
}

func (c *Calendar) generateSolsticeTable(latex string) string {
	const layout = "021504Z Jan"

	table := templates.SolsticeTableTemplate
	table = strings.Replace(table, "+FYS", strconv.Itoa(c.calendar.FiscalYear()-1), 1)
	table = strings.Replace(table, "+WS1", c.calendar.SolsticeTable().FirstWinterSolstice().Format(layout), 1)
	table = strings.Replace(table, "+VE", c.calendar.SolsticeTable().VernalEquinox().Format(layout), 1)
	table = strings.Replace(table, "+SS", c.calendar.SolsticeTable().SummerSolstice().Format(layout), 1)
	table = strings.Replace(table, "+AE", c.calendar.SolsticeTable().AutumnalEquinox().Format(layout), 1)
	table = strings.Replace(table, "+FYE", strconv.Itoa(c.calendar.FiscalYear()), 1)
	table = strings.Replace(table, "+WS2", c.calendar.SolsticeTable().SecondWinterSolstice().Format(layout), 1)

	latex = strings.Replace(latex, "+SOLSTICES", table, 1)

	return latex
}

func (c *Calendar) generateHolidayTables(latex string) string {
	ht := NewHolidayTables(c.calendar.HolidayCalendar(), c.calendar.FiscalYear())

	latex = ht.TableByOccurrence(latex)
	latex = ht.TableByAbbreviation(latex)

	return latex
}

func (c *Calendar) generateMiniMonthCmds(latex string) string {
	mm := ""
	for _, m := range c.minimonths {
		mm += m.LaTeX()
	}

	latex = strings.Replace(latex, "+MINIMONTH_CMDS", mm, 1)

	return latex
}

func (c *Calendar) generateTrimesterPages(latex string) string {
	trimester := c.calendar.FirstMonth().FirstWeek().Trimester()
	for i := 1; i <= trimesterCount; i++ {
		tr := NewTrimester(trimester, c.minimonths)

		latex = strings.Replace(latex, fmt.Sprintf("+T%d", i), tr.LaTeX(), 1)

		trimester = trimester.Next()
	}

	return latex
}

func (c *Calendar) generateQuarterPages(latex string) string {
	quarter := c.calendar.FirstMonth().FirstWeek().FiscalQuarter()
	for i := 1; i <= quarterCount; i++ {
		qt := NewQuarter(quarter, c.minimonths)

		latex = strings.Replace(latex, fmt.Sprintf("+T%d", i), qt.LaTeX(), 1)

		quarter = quarter.Next()
	}

	return latex
}

func (c *Calendar) generateMonthPages(latex string) string {
	month := c.calendar.FirstMonth()
	for i := 1; i <= calendarMonthCount; i++ {
		mo := NewMonth(c.calendar, month, c.minimonths)
		latex = strings.Replace(latex, fmt.Sprintf("+M%0d", i), mo.LaTeX(), 1)
		month = month.Next()
	}

	return latex
}

func (c *Calendar) LaTeX() string {
	latex := templates.CalendarTemplate

	latex = c.fillCalParams(latex)
	latex = c.generateDoomsdayTable(latex)
	latex = c.generateSolsticeTable(latex)
	latex = c.generateHolidayTables(latex)
	latex = c.generateMiniMonthCmds(latex)
	latex = c.generateTrimesterPages(latex)
	latex = c.generateQuarterPages(latex)
	latex = c.generateMonthPages(latex)

	return latex
}
