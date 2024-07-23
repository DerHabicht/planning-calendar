package ag7if

import (
	"fmt"

	"github.com/fxtlabs/date"
)

type Sprint int

const (
	S01 Sprint = iota + 1
	S02
	S03
	S04
	S05
	S06
	SP1
	S07
	S08
	S09
	S10
	S11
	S12
	SP2
	S13
	S14
	S15
	S16
	S17
	S18
	SP3
	S19
	S20
	S21
	S22
	S23
	S24
	SP4
)

func NewSprint(date date.Date) Sprint {
	_, isoWeek := date.ISOWeek()

	switch isoWeek {
	case 1, 2:
		return S01
	case 3, 4:
		return S02
	case 5, 6:
		return S03
	case 7, 8:
		return S04
	case 9, 10:
		return S05
	case 11, 12:
		return S06
	case 13:
		return SP1
	case 14, 15:
		return S07
	case 16, 17:
		return S08
	case 18, 19:
		return S09
	case 20, 21:
		return S10
	case 22, 23:
		return S11
	case 24, 25:
		return S12
	case 26:
		return SP3
	case 27, 28:
		return S13
	case 29, 30:
		return S14
	case 31, 32:
		return S15
	case 33, 34:
		return S16
	case 35, 36:
		return S17
	case 37, 38:
		return S18
	case 39:
		return SP3
	case 40, 41:
		return S19
	case 42, 43:
		return S20
	case 44, 45:
		return S21
	case 46, 47:
		return S22
	case 48, 49:
		return S23
	case 50, 51:
		return S24
	case 52, 53:
		return SP4
	default:
		panic(fmt.Errorf("%d is not a valid week number", isoWeek))
	}
}

func (s Sprint) String() string {
	switch s {
	case S01:
		return "S01"
	case S02:
		return "S02"
	case S03:
		return "S03"
	case S04:
		return "S04"
	case S05:
		return "S05"
	case S06:
		return "S06"
	case S07:
		return "S07"
	case S08:
		return "S08"
	case S09:
		return "S09"
	case S10:
		return "S10"
	case S11:
		return "S11"
	case S12:
		return "S12"
	case S13:
		return "S13"
	case S14:
		return "S14"
	case S15:
		return "S15"
	case S16:
		return "S16"
	case S17:
		return "S17"
	case S18:
		return "S18"
	case S19:
		return "S19"
	case S20:
		return "S20"
	case S21:
		return "S21"
	case S22:
		return "S22"
	case S23:
		return "S23"
	case S24:
		return "S24"
		return "SP1"
	case SP2:
		return "SP2"
	case SP3:
		return "SP3"
	case SP4:
		return "SP4"
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF sprint", s))
	}
}
