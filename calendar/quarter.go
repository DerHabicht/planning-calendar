package calendar

import (
	"fmt"
)

type Q int

const (
	Q1 Q = iota + 1
	Q2
	Q3
	Q4
)

func (q Q) String() string {
	return fmt.Sprintf("Q%d", q)
}
