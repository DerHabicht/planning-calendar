package calendar

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
	"github.com/rickar/cal/v2/ch"
	"github.com/rickar/cal/v2/de"
	"github.com/rickar/cal/v2/us"
)

var (
	ChiltonHawksBirthday = &cal.Holiday{
		Name:      "Chilton Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1955,
		Month:     time.January,
		Day:       30,
		Func:      cal.CalcDayOfMonth,
	}

	PalmSunday = &cal.Holiday{
		Name:   "Palm Sunday",
		Type:   cal.ObservanceReligious,
		Offset: -7,
		Func:   cal.CalcEasterOffset,
	}

	AnnualGeneralConferenceSunday = &cal.Holiday{
		Name:      "Annual General Conference Sunday",
		Type:      cal.ObservanceReligious,
		StartYear: 1830,
		Month:     time.April,
		Weekday:   time.Sunday,
		Offset:    1,
		Func:      cal.CalcWeekdayOffset,
	}

	AnnualGeneralConferenceSaturday = &cal.Holiday{
		Name:      "Annual General Conference Saturday",
		Type:      cal.ObservanceReligious,
		StartYear: 1830,
		Func: func(h *cal.Holiday, year int) time.Time {

			date := AnnualGeneralConferenceSunday.Func(AnnualGeneralConferenceSunday, year)

			return date.AddDate(0, 0, -1)
		},
	}

	ActualBirthOfChrist = &cal.Holiday{
		Name:  "Birth of Christ",
		Type:  cal.ObservanceReligious,
		Month: time.April,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	YurisNight = &cal.Holiday{
		Name:      "Yuri's Night",
		Type:      cal.ObservanceOther,
		StartYear: 1961,
		Month:     time.April,
		Day:       12,
		Func:      cal.CalcDayOfMonth,
	}

	RobinHawksBirthday = &cal.Holiday{
		Name:      "Robin Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1953,
		Month:     time.April,
		Day:       20,
		Func:      cal.CalcDayOfMonth,
	}

	RestorationOfThePriesthood = &cal.Holiday{
		Name:      "Restoration of the Priesthood",
		Type:      cal.ObservanceReligious,
		StartYear: 1829,
		Month:     time.May,
		Day:       15,
		Func:      cal.CalcDayOfMonth,
	}

	MoonLandingDay = &cal.Holiday{
		Name:      "Moon Landing Day",
		Type:      cal.ObservanceOther,
		StartYear: 1969,
		Month:     time.July,
		Day:       20,
		Func:      cal.CalcDayOfMonth,
	}

	PioneerDay = &cal.Holiday{
		Name:      "Pioneer Day",
		Type:      cal.ObservancePublic,
		StartYear: 1847,
		Month:     time.July,
		Day:       24,
		Func:      cal.CalcDayOfMonth,
	}

	RobertHawksBirthday = &cal.Holiday{
		Name:      "Robert Hawk's Birthday",
		Type:      cal.ObservanceOther,
		StartYear: 1988,
		Month:     time.September,
		Day:       27,
		Func:      cal.CalcDayOfMonth,
	}

	SemiAnnualGeneralConferenceSunday = &cal.Holiday{
		Name:      "Semi-Annual General Conference Sunday",
		Type:      cal.ObservanceReligious,
		StartYear: 1830,
		Month:     time.October,
		Weekday:   time.Sunday,
		Offset:    1,
		Func:      cal.CalcWeekdayOffset,
	}

	SemiAnnualGeneralConferenceSaturday = &cal.Holiday{
		Name:      "Semi-Annual General Conference Saturday",
		Type:      cal.ObservanceReligious,
		StartYear: 1830,
		Func: func(h *cal.Holiday, year int) time.Time {
			date := SemiAnnualGeneralConferenceSunday.Func(SemiAnnualGeneralConferenceSunday, year)

			return date.AddDate(0, 0, -1)
		},
	}

	CAPSunday = &cal.Holiday{
		Name:      "CAP Sunday",
		Type:      cal.ObservancePublic,
		StartYear: 1941,
		Month:     time.December,
		Weekday:   time.Sunday,
		Offset:    1,
		Func:      cal.CalcWeekdayOffset,
	}

	FirstFlightDay = &cal.Holiday{
		Name:      "First Flight Day",
		Type:      cal.ObservanceOther,
		StartYear: 1903,
		Month:     time.December,
		Day:       17,
		Func:      cal.CalcDayOfMonth,
	}

	FirstAdvent = &cal.Holiday{
		Name:    "First Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -4,
		Func:    cal.CalcWeekdayFrom,
	}

	SecondAdvent = &cal.Holiday{
		Name:    "Second Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -3,
		Func:    cal.CalcWeekdayFrom,
	}

	ThirdAdvent = &cal.Holiday{
		Name:    "Third Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -2,
		Func:    cal.CalcWeekdayFrom,
	}

	FourthAdvent = &cal.Holiday{
		Name:    "Fourth Advent",
		Type:    cal.ObservanceReligious,
		Month:   time.December,
		Day:     24,
		Weekday: time.Sunday,
		Offset:  -1,
		Func:    cal.CalcWeekdayFrom,
	}

	ChristmasEve = &cal.Holiday{
		Name:  "Christmas Eve",
		Type:  cal.ObservanceReligious,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	NewYearEve = &cal.Holiday{
		Name:  "New Year Eve",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	Ag7ifHolidays = map[string]*cal.Holiday{
		"NYD":   us.NewYear,
		"CHHBD": ChiltonHawksBirthday,
		"MLKD":  us.MlkDay,
		"PRESD": us.PresidentsDay,
		"PALMS": PalmSunday,
		"MAUNT": aa.MaundyThursday,
		"GOODF": aa.GoodFriday,
		"EASTR": aa.Easter,
		"JGKS":  AnnualGeneralConferenceSaturday,
		"JGKU":  AnnualGeneralConferenceSunday,
		"ABDOC": ActualBirthOfChrist,
		"YURIN": YurisNight,
		"RDHBD": RobinHawksBirthday,
		"APRD":  RestorationOfThePriesthood,
		"MEMOD": us.MemorialDay,
		"EMANC": us.Juneteenth,
		"INDEP": us.IndependenceDay,
		"MOONL": MoonLandingDay,
		"PIOND": PioneerDay,
		"CHBFT": ch.Bundesfeiertag.Clone(&cal.Holiday{Name: "Swiss National Holiday"}),
		"LABOR": us.LaborDay,
		"RHHBD": RobertHawksBirthday,
		"HJGKS": SemiAnnualGeneralConferenceSaturday,
		"HJGKU": SemiAnnualGeneralConferenceSunday,
		"DEEIN": de.DeutschenEinheit.Clone(&cal.Holiday{Name: "German Reunification Day"}),
		"CAPSD": CAPSunday,
		"COLUM": us.ColumbusDay,
		"VETD":  us.VeteransDay,
		"THXD":  us.ThanksgivingDay,
		"FFLTD": FirstFlightDay,
		"1ADV":  FirstAdvent,
		"2ADV":  SecondAdvent,
		"3ADV":  ThirdAdvent,
		"4ADV":  FourthAdvent,
		"XMASE": ChristmasEve,
		"XMASX": us.ChristmasDay,
		"STEPH": aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Feast of St. Stephen"}),
		"NYE":   NewYearEve,
	}
)
