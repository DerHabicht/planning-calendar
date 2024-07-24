package calendar

import (
	"fmt"
)

type T int

const (
	T1 T = iota + 1
	T2
	T3
)

func (t T) String() string {
	return fmt.Sprintf("T%d", t)
}
