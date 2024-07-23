package calendar

type Solstice int

const (
	NoSolstice Solstice = iota
	VernalEquinox
	SummerSolstice
	AutumnalEquinox
	WinterSolstice
)

func (s Solstice) LaTeX() string {
	switch s {
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
