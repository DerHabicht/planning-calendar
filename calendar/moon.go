package calendar

import (
	"strings"

	"github.com/pkg/errors"
)

type MoonPhase int

const (
	NewMoon MoonPhase = iota
	WaxingCrescent
	FirstQuarter
	WaxingGibbous
	FullMoon
	WaningGibbous
	ThirdQuarter
	WaningCrescent
)

func ParseMoonPhase(s string) (MoonPhase, error) {
	switch strings.ToLower(s) {
	case "new moon":
		return NewMoon, nil
	case "waxing crescent":
		return WaxingCrescent, nil
	case "first quarter":
		return FirstQuarter, nil
	case "waxing gibbous":
		return WaxingGibbous, nil
	case "full moon":
		return FullMoon, nil
	case "waning gibbous":
		return WaningGibbous, nil
	case "third quarter":
		return ThirdQuarter, nil
	case "waning crescent":
		return WaningCrescent, nil
	default:
		return -1, errors.Errorf("unrecognized moon phase: %s", s)
	}
}

func (mp MoonPhase) String() string {
	switch mp {
	case NewMoon:
		return "New Moon"
	case WaxingCrescent:
		return "Waxing Crescent"
	case FirstQuarter:
		return "First Quarter"
	case WaxingGibbous:
		return "Waxing Gibbous"
	case FullMoon:
		return "Full Moon"
	case WaningGibbous:
		return "Waning Gibbous"
	case ThirdQuarter:
		return "Third Quarter"
	case WaningCrescent:
		return "Waning Crescent"
	default:
		panic(errors.Errorf("invalid MoonPhase value: %d", mp))
	}
}

func (mp MoonPhase) LaTeX() string {
	switch mp {
	case NewMoon:
		return `\newmoon{}`
	case WaxingCrescent:
		return `\waxingcrescent{}`
	case FirstQuarter:
		return `\firstquartermoon{}`
	case WaxingGibbous:
		return `\waxinggibbous{}`
	case FullMoon:
		return `\fullmoon{}`
	case WaningGibbous:
		return `\waninggibbous{}`
	case ThirdQuarter:
		return `\thirdquartermoon{}`
	case WaningCrescent:
		return `\waningcrescent{}`
	default:
		panic(errors.Errorf("invalid MoonPhase value: %d", mp))
	}
}
