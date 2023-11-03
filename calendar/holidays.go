package calendar

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
	"github.com/rickar/cal/v2/ch"
	"github.com/rickar/cal/v2/de"
	"github.com/rickar/cal/v2/us"
)

// TODO: Break this out into a config file

var Ag7ifHolidays map[string]*cal.Holiday = map[string]*cal.Holiday{
	"NYD": us.NewYear,
	"CHHBD": {
		Name:      "Chilton Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1955,
		Month:     time.January,
		Day:       30,
		Func:      cal.CalcDayOfMonth,
	},
	"MLKD":  us.MlkDay,
	"PRESD": us.PresidentsDay,
	"PALMS": {
		Name:   "Palm Sunday",
		Type:   cal.ObservanceReligious,
		Offset: -7,
		Func:   cal.CalcEasterOffset,
	},
	"MAUNT": aa.MaundyThursday,
	"GOODF": aa.GoodFriday,
	"EASTR": aa.Easter,
	"ABDOC": {
		Name:  "Birth of Christ",
		Type:  cal.ObservanceReligious,
		Month: time.April,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	},
	"YURIN": {
		Name:      "Yuri's Night",
		Type:      cal.ObservanceOther,
		StartYear: 1961,
		Month:     time.April,
		Day:       12,
		Func:      cal.CalcDayOfMonth,
	},
	"RDHBD": {
		Name:      "Robin Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1953,
		Month:     time.April,
		Day:       20,
		Func:      cal.CalcDayOfMonth,
	},
	"APRD": {
		Name:      "Restoration of the Priesthood",
		Type:      cal.ObservanceReligious,
		StartYear: 1829,
		Month:     time.May,
		Day:       15,
		Func:      cal.CalcDayOfMonth,
	},
	"MEMOD": us.MemorialDay,
	"EMANC": us.Juneteenth,
	"INDEP": us.IndependenceDay,
	"MOONL": {
		Name:      "Moon Landing Day",
		Type:      cal.ObservanceOther,
		StartYear: 1969,
		Month:     time.July,
		Day:       20,
		Func:      cal.CalcDayOfMonth,
	},
	"PIOND": {
		Name:      "Pioneer Day",
		Type:      cal.ObservancePublic,
		StartYear: 1847,
		Month:     time.July,
		Day:       24,
		Func:      cal.CalcDayOfMonth,
	},
	"CHBFT": ch.Bundesfeiertag.Clone(&cal.Holiday{Name: "Swiss National Holiday"}),
	"LABOR": us.LaborDay,
	"RHHBD": {
		Name:      "Robert Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1988,
		Month:     time.September,
		Day:       27,
		Func:      cal.CalcDayOfMonth,
	},
	"DEEIN": de.DeutschenEinheit.Clone(&cal.Holiday{Name: "German Reunification Day"}),
	"CAPSD": {
		Name:      "CAP Sunday",
		Type:      cal.ObservancePublic,
		StartYear: 1941,
		Month:     time.December,
		Weekday:   time.Sunday,
		Offset:    1,
		Func:      cal.CalcWeekdayOffset,
	},
	"COLUM": us.ColumbusDay,
	"VETD":  us.VeteransDay,
	"THXD":  us.ThanksgivingDay,
	"FFLTD": {
		Name:      "First Flight Day",
		Type:      cal.ObservanceOther,
		StartYear: 1903,
		Month:     time.December,
		Day:       17,
		Func:      cal.CalcDayOfMonth,
	},
	"1ADV": {
		Name:    "First Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -4,
		Func:    cal.CalcWeekdayFrom,
	},
	"2ADV": {
		Name:    "Second Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -3,
		Func:    cal.CalcWeekdayFrom,
	},
	"3ADV": {
		Name:    "Third Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -2,
		Func:    cal.CalcWeekdayFrom,
	},
	"4ADV": {
		Name:    "Fourth Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -1,
		Func:    cal.CalcWeekdayFrom,
	},
	"XMASE": {
		Name:  "Christmas Eve",
		Type:  cal.ObservanceReligious,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	},
	"XMASX": us.ChristmasDay,
	"STEPH": aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Feast of St. Stephen"}),
	"NYE": {
		Name:  "New Year's Eve",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	},
}
