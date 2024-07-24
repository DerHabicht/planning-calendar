package ag7if

import (
	"fmt"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

type S int

const (
	S01 S = iota + 1
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

func ComputeSprint(date date.Date) (int, S) {
	year, isoWeek := date.ISOWeek()

	switch isoWeek {
	case 1, 2:
		return year, S01
	case 3, 4:
		return year, S02
	case 5, 6:
		return year, S03
	case 7, 8:
		return year, S04
	case 9, 10:
		return year, S05
	case 11, 12:
		return year, S06
	case 13:
		return year, SP1
	case 14, 15:
		return year, S07
	case 16, 17:
		return year, S08
	case 18, 19:
		return year, S09
	case 20, 21:
		return year, S10
	case 22, 23:
		return year, S11
	case 24, 25:
		return year, S12
	case 26:
		return year, SP3
	case 27, 28:
		return year, S13
	case 29, 30:
		return year, S14
	case 31, 32:
		return year, S15
	case 33, 34:
		return year, S16
	case 35, 36:
		return year, S17
	case 37, 38:
		return year, S18
	case 39:
		return year, SP3
	case 40, 41:
		return year, S19
	case 42, 43:
		return year, S20
	case 44, 45:
		return year, S21
	case 46, 47:
		return year, S22
	case 48, 49:
		return year, S23
	case 50, 51:
		return year, S24
	case 52, 53:
		return year, SP4
	default:
		panic(fmt.Errorf("%d is not a valid week number", isoWeek))
	}
}

func ComputeWeekNumbers(s S) []int {
	switch s {
	case S01:
		return []int{1, 2}
	case S02:
		return []int{3, 4}
	case S03:
		return []int{5, 6}
	case S04:
		return []int{7, 8}
	case S05:
		return []int{9, 10}
	case S06:
		return []int{11, 12}
	case SP1:
		return []int{13}
	case S07:
		return []int{14, 15}
	case S08:
		return []int{16, 17}
	case S09:
		return []int{18, 19}
	case S10:
		return []int{20, 21}
	case S11:
		return []int{22, 23}
	case S12:
		return []int{24, 25}
	case SP2:
		return []int{26}
	case S13:
		return []int{27, 28}
	case S14:
		return []int{29, 30}
	case S15:
		return []int{31, 32}
	case S16:
		return []int{33, 34}
	case S17:
		return []int{35, 36}
	case S18:
		return []int{37, 38}
	case SP3:
		return []int{39}
	case S19:
		return []int{40, 41}
	case S20:
		return []int{42, 43}
	case S21:
		return []int{44, 45}
	case S22:
		return []int{46, 47}
	case S23:
		return []int{48, 49}
	case S24:
		return []int{50, 51}
	case SP4:
		return []int{52, 53}
	default:
		panic(errors.Errorf("invalid value for sprint: %d", s))
	}
}

func (s S) String() string {
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
	case SP1:
		return "SP1"
	case SP2:
		return "SP2"
	case SP3:
		return "SP3"
	case SP4:
		return "SP4"
	default:
		panic(fmt.Errorf("%d is not a valid sprint", s))
	}
}
