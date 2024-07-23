package events

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"

	"github.com/derhabicht/planning-calendar/calendar"
)

func TimeToLocalDate(t time.Time) date.Date {
	local := t.In(time.Local)

	return date.New(local.Year(), local.Month(), local.Day())
}

type SolsticeTable struct {
	fiscalYear           int
	firstWinterSolstice  time.Time
	vernalEquinox        time.Time
	summerSolstice       time.Time
	autumnalEquinox      time.Time
	secondWinterSolstice time.Time
}

func NewSolsticeTable(fy int) SolsticeTable {
	vernal := julian.JDToTime(solstice.March(fy))
	summer := julian.JDToTime(solstice.June(fy))
	autumnal := julian.JDToTime(solstice.September(fy))
	winter1 := julian.JDToTime(solstice.December(fy - 1))
	winter2 := julian.JDToTime(solstice.December(fy))

	return SolsticeTable{
		fiscalYear:           fy,
		firstWinterSolstice:  winter1,
		vernalEquinox:        vernal,
		summerSolstice:       summer,
		autumnalEquinox:      autumnal,
		secondWinterSolstice: winter2,
	}
}

func (st SolsticeTable) FiscalYear() int {
	return st.fiscalYear
}

func (st SolsticeTable) FirstWinterSolstice() time.Time {
	return st.firstWinterSolstice
}

func (st SolsticeTable) VernalEquinox() time.Time {
	return st.vernalEquinox
}

func (st SolsticeTable) SummerSolstice() time.Time {
	return st.summerSolstice
}

func (st SolsticeTable) AutumnalEquinox() time.Time {
	return st.autumnalEquinox
}

func (st SolsticeTable) SecondWinterSolstice() time.Time {
	return st.secondWinterSolstice
}

func (st SolsticeTable) IsSolstice(d date.Date) (bool, calendar.Solstice) {
	switch d.Month() {
	case time.March:
		if d == TimeToLocalDate(st.vernalEquinox) {
			return true, calendar.VernalEquinox
		}
	case time.June:
		if d == TimeToLocalDate(st.summerSolstice) {
			return true, calendar.SummerSolstice
		}
	case time.September:
		if d == TimeToLocalDate(st.autumnalEquinox) {
			return true, calendar.AutumnalEquinox
		}
	case time.December:
		if (d == TimeToLocalDate(st.firstWinterSolstice)) || (d == TimeToLocalDate(st.secondWinterSolstice)) {
			return true, calendar.WinterSolstice
		}
	default:
		break
	}

	return false, 0
}
