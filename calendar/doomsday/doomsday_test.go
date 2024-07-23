package doomsday

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestComputeDoomsday(t *testing.T) {
	testYears := map[int]time.Weekday{
		1700: time.Sunday,
		1800: time.Friday,
		1900: time.Wednesday,
		2000: time.Tuesday,
		2022: time.Monday,
		2023: time.Tuesday,
		2024: time.Thursday,
		2025: time.Friday,
		2026: time.Saturday,
	}

	for yr, weekday := range testYears {
		// Validate test cases (4 July lands on Doomsday)
		d := date.New(yr, time.July, 4)
		assert.Equalf(t, d.Weekday(), weekday, "Test case %d: %s is invalid.", yr, weekday)

		// Validate the algorithm
		res := ComputeDoomsday(yr)
		assert.Equalf(t, weekday, res, "Returned %s for %d, correct doomsday is %s", res, yr, weekday)
	}
}
