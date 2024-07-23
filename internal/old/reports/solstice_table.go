package reports

import (
	"fmt"
	"strings"
	"time"

	"github.com/derhabicht/planning-calendar/calendar/events"
)

type SolsticeSymbol int

const (
	VernalEquinox SolsticeSymbol = iota + 1
	SummerSolstice
	AutumnalEquinox
	WinterSolstice
)

func SolsticeSymbolFromSolstice(sol events.Solstice) SolsticeSymbol {
	switch sol {
	case events.VernalEquinox:
		return VernalEquinox
	case events.SummerSolstice:
		return SummerSolstice
	case events.AutumnalEquinox:
		return AutumnalEquinox
	case events.WinterSolstice:
		return WinterSolstice
	default:
		return 0
	}
}

func (ss SolsticeSymbol) LaTeX() string {
	switch ss {
	case VernalEquinox:
		return `\Aries`
	case SummerSolstice:
		return `\Cancer`
	case AutumnalEquinox:
		return `\Libra`
	case WinterSolstice:
		return `\Capricorn`
	default:
		return ""
	}
}

const solsticeTableTemplate string = `\begin{tabular}{llr}
\toprule
\Capricorn  & Winter Solstice (+%FYS) & +%WS1 \\
\Aries      & Vernal Equinox          & +%VE  \\
\Cancer     & Summer Solstice         & +%SS  \\
\Libra      & Autumnal Equinox        & +%AE  \\
\Capricorn  & Winter Solstice (+%FYE) & +%WS2 \\
\bottomrule
\end{tabular}
`

type SolsticeTable struct {
	fiscalYear           int
	firstWinterSolstice  time.Time
	vernalEquinox        time.Time
	summerSolstice       time.Time
	autumnalEquinox      time.Time
	secondWinterSolstice time.Time
}

func NewSolsticeTable(table events.SolsticeTable) SolsticeTable {
	return SolsticeTable{
		fiscalYear:           table.FiscalYear(),
		firstWinterSolstice:  table.FirstWinterSolstice(),
		vernalEquinox:        table.VernalEquinox(),
		summerSolstice:       table.SummerSolstice(),
		autumnalEquinox:      table.AutumnalEquinox(),
		secondWinterSolstice: table.SecondWinterSolstice(),
	}
}

func (st SolsticeTable) LaTeX() string {
	const layout = "021504Z Jan"
	latex := solsticeTableTemplate

	latex = strings.Replace(latex, "+%FYS", fmt.Sprintf("%d", st.fiscalYear-1), 1)
	latex = strings.Replace(latex, "+%FYE", fmt.Sprintf("%d", st.fiscalYear), 1)
	latex = strings.Replace(latex, "+%WS1", st.firstWinterSolstice.Format(layout), 1)
	latex = strings.Replace(latex, "+%VE", st.vernalEquinox.Format(layout), 1)
	latex = strings.Replace(latex, "+%SS", st.summerSolstice.Format(layout), 1)
	latex = strings.Replace(latex, "+%AE", st.autumnalEquinox.Format(layout), 1)
	latex = strings.Replace(latex, "+%WS2", st.secondWinterSolstice.Format(layout), 1)

	return latex
}
