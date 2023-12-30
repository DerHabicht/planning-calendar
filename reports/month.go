package reports

import (
	"fmt"
	"strings"
	"time"

	"github.com/fxtlabs/date"
	"github.com/rickar/cal/v2"

	"github.com/derhabicht/planning-calendar/calendar"
)

const dayDataTemplate = `+DY\\+HD\moon{+FD}\\\vspace{1em}\hspace{1em}+YD\hfill{}+SR\\+MJD\hfill{}+SS`

type MonthTemplate struct {
	calendar calendar.Calendar
	month    string
	template string
}

func NewMonthTemplate(cal calendar.Calendar, template string) MonthTemplate {
	return MonthTemplate{
		calendar: cal,
		month:    cal.CurrentMonth(),
		template: template,
	}
}

func (m *MonthTemplate) fillDayData(dayStr string, sunrise, sunset time.Time, dt date.Date) string {
	mjd := int(cal.ModifiedJulianDate(dt.Local()))

	data := strings.Replace(dayDataTemplate, "+DY", dayStr, 1)
	data = strings.Replace(data, "+SR", sunrise.Format("1504"), 1)
	data = strings.Replace(data, "+SS", sunset.Format("1504"), 1)
	data = strings.Replace(data, "+FD", dt.FormatISO(4), 1)
	data = strings.Replace(data, "+YD", fmt.Sprintf("%03d", dt.YearDay()), 1)
	data = strings.Replace(data, "+MJD", fmt.Sprintf("%d", mjd), 1)

	return data
}

func (m *MonthTemplate) fillHolidayData(dayStr string, holiday *calendar.Holiday, actual bool) string {
	data := dayStr

	if holiday == nil {
		data = strings.Replace(data, "+HD", "", 1)
	} else {
		abbv := holiday.Abbreviation()
		if !actual {
			abbv += "*"
		}
		data = strings.Replace(data, "+HD", fmt.Sprintf("%s\\hfill{}", abbv), 1)
	}

	return data
}

func (m *MonthTemplate) fillWeekData(weekNum int, week calendar.Week) {
	isoWeekString, weekCard := week.IsoWeek()
	isoWeekString = fmt.Sprintf("%s%s", MarshallCard(weekCard), isoWeekString)

	m.template = strings.Replace(m.template, fmt.Sprintf("+FT%d", weekNum), week.FyTrimester(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+FQ%d", weekNum), week.FyQuarter(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+FW%d", weekNum), week.FyWeek(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+AQ%d", weekNum), week.Ag7ifQuarter(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+AS%d", weekNum), week.Ag7ifSprint(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+IW%d", weekNum), isoWeekString, 1)
}

func (m *MonthTemplate) LaTeX() string {
	m.template = strings.Replace(m.template, "+M", m.calendar.CurrentMonth(), 1)

	day := 0
	week := m.calendar.CurrentWeek()

	for wk := 1; wk <= 6; wk++ {
		week.Reset()
		m.fillWeekData(wk, week)

		for i := 0; i < 7; i++ {
			var dayStr string
			var err error
			if day == 0 {
				dayStr, err = week.CurrentDayStr(true)
			} else {
				dayStr, err = week.CurrentDayStr(false)
			}
			if err != nil {
				wk, _ := week.IsoWeek()
				panic(fmt.Errorf("unexpectedly got to the end of week %s", wk))
			}

			currentDate, _ := week.CurrentDay()

			var holiday *calendar.Holiday
			actual, _, holiday := week.IsCurrentDayHoliday()

			sunrise, sunset, err := week.CurrentDaySunriseSunset()
			if err != nil {
				panic(err)
			}

			dayData := m.fillDayData(dayStr, sunrise, sunset, currentDate)
			dayData = m.fillHolidayData(dayData, holiday, actual)

			day += 1
			m.template = strings.Replace(m.template, fmt.Sprintf("+D%02d", day), dayData, 1)

			week.Next()
		}

		week = m.calendar.NextWeek()
	}

	return m.template
}
